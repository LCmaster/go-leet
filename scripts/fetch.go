// Package main provides a command-line tool to fetch LeetCode problems and
// automatically scaffold the solution folder, README, starter code, and tests.
//
// Usage:
//
//	go run ./scripts/fetch.go <problem_number_or_slug>
//
// Examples:
//
//	go run ./scripts/fetch.go 5
//	go run ./scripts/fetch.go longest-palindromic-substring
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ProblemsAllResponse models the JSON payload returned by LeetCode's public
// problem list REST API at https://leetcode.com/api/problems/all/.
type ProblemsAllResponse struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionID         int    `json:"question_id"`
			FrontendQuestionID any    `json:"frontend_question_id"` // Can be float64 or string
			QuestionTitle      string `json:"question__title"`
			QuestionTitleSlug  string `json:"question__title_slug"`
		} `json:"stat"`
	} `json:"stat_status_pairs"`
}

// GraphQLResponse models the JSON response from LeetCode's GraphQL API
// for the "questionData" query.
type GraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionFrontendID string `json:"questionFrontendId"`
			Title              string `json:"title"`
			TitleSlug          string `json:"titleSlug"`
			Difficulty         string `json:"difficulty"`
			Content            string `json:"content"`
			CodeSnippets       []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
			} `json:"codeSnippets"`
		} `json:"question"`
	} `json:"data"`
}

// QuestionData holds normalized problem details extracted from the GraphQL query.
type QuestionData struct {
	QuestionFrontendID string
	Title              string
	TitleSlug          string
	Difficulty         string
	Content            string
	CodeSnippets       []struct {
		Lang     string `json:"lang"`
		LangSlug string `json:"langSlug"`
		Code     string `json:"code"`
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./scripts/fetch.go <problem_number_or_slug>")
		os.Exit(1)
	}

	target := strings.TrimSpace(os.Args[1])

	// 1. Resolve problem slug from input (e.g. "5" -> "longest-palindromic-substring")
	slug, err := resolveSlug(target)
	if err != nil {
		fmt.Printf("Error resolving problem '%s': %v\n", target, err)
		os.Exit(1)
	}

	// 2. Fetch problem details and Go code snippet via LeetCode GraphQL
	q, err := fetchQuestion(slug)
	if err != nil {
		fmt.Printf("Error fetching question '%s': %v\n", slug, err)
		os.Exit(1)
	}

	// 3. Determine directory name following the format: 0005_longest_palindromic_substring
	idNum, _ := strconv.Atoi(q.QuestionFrontendID)
	folderSnake := strings.ReplaceAll(q.TitleSlug, "-", "_")
	dirName := fmt.Sprintf("%04d_%s", idNum, folderSnake)
	targetDir := filepath.Join(".", dirName)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", targetDir, err)
		os.Exit(1)
	}

	// 4. Create README.md containing the formatted problem description
	readmePath := filepath.Join(targetDir, "README.md")
	readmeContent := buildReadme(q.QuestionFrontendID, q.Title, q.Difficulty, q.Content)
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		fmt.Printf("Error writing README.md: %v\n", err)
	} else {
		fmt.Printf("Created/Updated %s\n", readmePath)
	}

	// 5. Create solution Go file with starter template if absent or empty
	goFileName := folderSnake + ".go"
	goFilePath := filepath.Join(targetDir, goFileName)
	pkgName := strings.ReplaceAll(folderSnake, "_", "")

	goCodeSnippet := ""
	for _, snippet := range q.CodeSnippets {
		if snippet.LangSlug == "golang" {
			goCodeSnippet = snippet.Code
			break
		}
	}

	if shouldWriteFile(goFilePath) {
		commentBlock := buildCommentBlock(q.QuestionFrontendID, q.Title, q.Content)
		starterCode := fmt.Sprintf("package %s\n\n%s\n\n%s\n", pkgName, commentBlock, goCodeSnippet)
		if err := os.WriteFile(goFilePath, []byte(starterCode), 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", goFilePath, err)
		} else {
			fmt.Printf("Created %s\n", goFilePath)
		}
	} else {
		fmt.Printf("Skipped %s (already exists and has content)\n", goFilePath)
	}

	// 6. Create unit test Go file template if absent or empty
	testFileName := folderSnake + "_test.go"
	testFilePath := filepath.Join(targetDir, testFileName)
	if shouldWriteFile(testFilePath) {
		testStarter := fmt.Sprintf("package %s\n\nimport (\n\t\"testing\"\n)\n\nfunc TestSolution(t *testing.T) {\n\t// TODO: add tests\n}\n", pkgName)
		if err := os.WriteFile(testFilePath, []byte(testStarter), 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", testFilePath, err)
		} else {
			fmt.Printf("Created %s\n", testFilePath)
		}
	}

	fmt.Printf("\nDone! Problem %s (%s) set up in %s\n", q.QuestionFrontendID, q.Title, targetDir)
}

// shouldWriteFile checks whether the target file does not exist or has 0 bytes.
// This ensures that existing solutions are never overwritten accidentally.
func shouldWriteFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	return info.Size() == 0
}

// resolveSlug converts numeric problem IDs (e.g. "5", "0005") to their slug equivalent
// by querying LeetCode's public problem listing API. If the input is already a slug,
// it sanitizes and returns it directly.
func resolveSlug(target string) (string, error) {
	// If target is already a slug (contains non-numeric characters), sanitize and return.
	if _, err := strconv.Atoi(target); err != nil {
		return strings.ReplaceAll(target, "_", "-"), nil
	}

	targetNum, _ := strconv.Atoi(target)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://leetcode.com/api/problems/all/", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data ProblemsAllResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	for _, pair := range data.StatStatusPairs {
		var fID int
		switch v := pair.Stat.FrontendQuestionID.(type) {
		case float64:
			fID = int(v)
		case string:
			fID, _ = strconv.Atoi(v)
		}

		if fID == targetNum {
			return pair.Stat.QuestionTitleSlug, nil
		}
	}

	return "", fmt.Errorf("problem #%d not found", targetNum)
}

// fetchQuestion queries LeetCode's official GraphQL endpoint for question details,
// difficulty, HTML content, and official code snippets.
func fetchQuestion(slug string) (*QuestionData, error) {
	query := `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId
    title
    titleSlug
    difficulty
    content
    codeSnippets {
      lang
      langSlug
      code
    }
  }
}`

	reqBody := map[string]any{
		"operationName": "questionData",
		"query":         query,
		"variables": map[string]string{
			"titleSlug": slug,
		},
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gqlResp GraphQLResponse
	if err := json.Unmarshal(respBytes, &gqlResp); err != nil {
		return nil, err
	}

	q := gqlResp.Data.Question
	if q.TitleSlug == "" {
		return nil, fmt.Errorf("empty question response for slug '%s'", slug)
	}

	return &QuestionData{
		QuestionFrontendID: q.QuestionFrontendID,
		Title:              q.Title,
		TitleSlug:          q.TitleSlug,
		Difficulty:         q.Difficulty,
		Content:            q.Content,
		CodeSnippets:       q.CodeSnippets,
	}, nil
}

// buildReadme formats the problem into a Markdown README structure.
func buildReadme(id, title, difficulty, htmlContent string) string {
	mdContent := htmlToMarkdown(htmlContent)
	return fmt.Sprintf("# %s. %s\n\n**Difficulty:** %s\n\n%s\n", id, title, difficulty, mdContent)
}

// buildCommentBlock formats the problem description using "//" line comments
// for every line (including empty lines), adhering to Go documentation conventions.
func buildCommentBlock(id, title, htmlContent string) string {
	mdContent := htmlToMarkdown(htmlContent)
	fullText := fmt.Sprintf("%s. %s\n\n%s", id, title, mdContent)

	lines := strings.Split(fullText, "\n")
	var commented []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			commented = append(commented, "//")
		} else {
			commented = append(commented, "// "+line)
		}
	}
	return strings.Join(commented, "\n")
}

// htmlToMarkdown parses and cleans LeetCode's raw HTML problem description,
// converting tags (<pre>, <code>, <strong>, <li>, <p>) into GitHub Flavored Markdown.
func htmlToMarkdown(htmlStr string) string {
	s := htmlStr

	// Replace preformatted example/code blocks (<pre>...</pre>)
	reCodeBlock := regexp.MustCompile(`(?s)<pre>(.*?)</pre>`)
	s = reCodeBlock.ReplaceAllStringFunc(s, func(match string) string {
		inner := reCodeBlock.FindStringSubmatch(match)[1]
		cleanInner := regexp.MustCompile(`<[^>]+>`).ReplaceAllString(inner, "")
		cleanInner = html.UnescapeString(cleanInner)
		cleanInner = strings.TrimSpace(cleanInner)
		return "\n```text\n" + cleanInner + "\n```\n"
	})

	// Common inline formatting
	s = regexp.MustCompile(`<strong>(.*?)</strong>`).ReplaceAllString(s, "**$1**")
	s = regexp.MustCompile(`<b>(.*?)</b>`).ReplaceAllString(s, "**$1**")
	s = regexp.MustCompile(`<em>(.*?)</em>`).ReplaceAllString(s, "*$1*")
	s = regexp.MustCompile(`<i>(.*?)</i>`).ReplaceAllString(s, "*$1*")
	s = regexp.MustCompile(`<code>(.*?)</code>`).ReplaceAllString(s, "`$1`")

	// Paragraphs & line breaks
	s = strings.ReplaceAll(s, "<p>", "\n")
	s = strings.ReplaceAll(s, "</p>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br>", "\n")

	// Strip non-breaking spaces and HTML noise
	s = strings.ReplaceAll(s, "\u00a0", "")
	s = strings.ReplaceAll(s, "&nbsp;", "")

	// Unordered list items
	s = regexp.MustCompile(`<li>\s*(.*?)\s*</li>`).ReplaceAllString(s, "- $1\n")
	s = strings.ReplaceAll(s, "<ul>", "\n")
	s = strings.ReplaceAll(s, "</ul>", "\n")
	s = strings.ReplaceAll(s, "<ol>", "\n")
	s = strings.ReplaceAll(s, "</ol>", "\n")

	// Strip any remaining HTML tags
	s = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(s, "")

	// Decode leftover HTML entities (&lt;, &gt;, etc.)
	s = html.UnescapeString(s)

	// Normalize list item indents
	s = regexp.MustCompile(`(?m)^\s+-\s+`).ReplaceAllString(s, "- ")

	// Clean up excessive blank lines
	reMultipleNewlines := regexp.MustCompile(`\n{3,}`)
	s = reMultipleNewlines.ReplaceAllString(s, "\n\n")

	return strings.TrimSpace(s)
}
