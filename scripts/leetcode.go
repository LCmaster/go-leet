// Package main is a single command-line tool for working with LeetCode problems
// in this repository: scaffold a problem locally, run it against LeetCode's
// sample cases, and submit it for full judging.
//
// Usage:
//
//	go run ./scripts/leetcode.go [-fetch] [-test] [-submit] [-input "<cases>"] [-input-file <file>] <problem>
//
// <problem> can be a problem number (15), a slug (3sum) or a path to the
// problem folder/file (0010-0019/0015_threesum). Flags may appear before or
// after the problem. When several actions are given they run in the order
// fetch → test → submit, and submit is skipped if the test fails.
//
// Examples:
//
//	go run ./scripts/leetcode.go -fetch 16
//	go run ./scripts/leetcode.go -test 16
//	go run ./scripts/leetcode.go -test 16 -input "[-1,2,1,-4]\n1"
//	go run ./scripts/leetcode.go -submit 16
//	go run ./scripts/leetcode.go -test -submit 16
//
// -test and -submit authenticate with browser session cookies read from a
// .env file at the repository root (see .env.example):
//
//	LEETCODE_SESSION=<value of the LEETCODE_SESSION cookie>
//	LEETCODE_CSRFTOKEN=<value of the csrftoken cookie>
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL      = "https://leetcode.com"
	userAgent    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0 Safari/537.36"
	pollInterval = 1500 * time.Millisecond
	pollTimeout  = 90 * time.Second
)

// predefinedTypes are supplied by LeetCode's Go judge. Local declarations of
// these types (kept so the package compiles for `go test`) must be stripped
// before upload, otherwise LeetCode reports "ListNode redeclared".
var predefinedTypes = map[string]bool{
	"ListNode": true,
	"TreeNode": true,
	"Node":     true,
}

// ---------------------------------------------------------------------------
// Terminal styling
// ---------------------------------------------------------------------------

var useColor = os.Getenv("NO_COLOR") == ""

func color(code, s string) string {
	if !useColor {
		return s
	}
	return "\033[" + code + "m" + s + "\033[0m"
}

func green(s string) string  { return color("32", s) }
func red(s string) string    { return color("31", s) }
func yellow(s string) string { return color("33", s) }
func cyan(s string) string   { return color("36", s) }
func bold(s string) string   { return color("1", s) }
func dim(s string) string    { return color("2", s) }

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, red("error: ")+format+"\n", args...)
	os.Exit(1)
}

// ---------------------------------------------------------------------------
// Entry point
// ---------------------------------------------------------------------------

func main() {
	doFetch := flag.Bool("fetch", false, "scaffold the problem folder, README, starter code and test file")
	doTest := flag.Bool("test", false, "run the local solution against sample (or custom) test cases")
	doSubmit := flag.Bool("submit", false, "submit the local solution for full judging")
	input := flag.String("input", "", `custom test input for -test; use "\n" to separate lines`)
	inputFile := flag.String("input-file", "", "read custom test input for -test from a file")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: go run ./scripts/leetcode.go [-fetch] [-test] [-submit] [options] <problem_number|slug|path>")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
	}

	target, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, red("error: ")+err.Error())
		flag.Usage()
		os.Exit(2)
	}
	if !*doFetch && !*doTest && !*doSubmit {
		fmt.Fprintln(os.Stderr, red("error: ")+"specify at least one of -fetch, -test, -submit")
		flag.Usage()
		os.Exit(2)
	}

	root, err := findRepoRoot()
	if err != nil {
		fatalf("%v", err)
	}
	loadDotEnv(filepath.Join(root, ".env"))
	client := newLeetClient()

	ref, err := resolveTarget(target)
	if err != nil {
		fatalf("%v", err)
	}

	q, err := client.fetchQuestion(ref.Slug)
	if err != nil {
		fatalf("fetching question %q: %v", ref.Slug, err)
	}
	fmt.Printf("%s %s. %s %s\n", dim("•"), q.FrontendID, bold(q.Title), dim("("+q.Difficulty+")"))

	if *doFetch {
		if err := scaffold(root, q); err != nil {
			fatalf("%v", err)
		}
	}

	if !*doTest && !*doSubmit {
		return
	}

	username, err := client.checkAuth()
	if err != nil {
		fatalf("%v", err)
	}
	fmt.Printf("%s signed in as %s\n", dim("•"), bold(username))

	file := ref.File
	if file == "" {
		dir := ref.Dir
		if dir == "" {
			if dir, err = findProblemDir(root, ref.Num); err != nil {
				fatalf("%v", err)
			}
		}
		if file, err = pickSolutionFile(dir); err != nil {
			fatalf("%v", err)
		}
	}
	rel, _ := filepath.Rel(root, file)
	fmt.Printf("%s file: %s\n\n", dim("•"), filepath.ToSlash(rel))

	code, err := prepareCode(file)
	if err != nil {
		fatalf("preparing %s: %v", rel, err)
	}

	if *doTest {
		dataInput, cases, err := buildTestInput(q, *input, *inputFile)
		if err != nil {
			fatalf("%v", err)
		}
		if !runTest(client, q, code, dataInput, cases) {
			if *doSubmit {
				fmt.Println(yellow("\nSkipping submit because the test run failed."))
			}
			os.Exit(1)
		}
		if *doSubmit {
			fmt.Println()
		}
	}

	if *doSubmit && !runSubmit(client, q, code) {
		os.Exit(1)
	}
}

// parseArgs parses flags that may appear before or after the single
// positional <problem> argument (the stdlib flag package stops at the first
// positional argument, so parsing is resumed after it).
func parseArgs(args []string) (string, error) {
	var positional []string
	for {
		if err := flag.CommandLine.Parse(args); err != nil {
			return "", err
		}
		if flag.NArg() == 0 {
			break
		}
		positional = append(positional, flag.Arg(0))
		args = flag.Args()[1:]
	}
	switch len(positional) {
	case 0:
		return "", errors.New("missing <problem> argument")
	case 1:
		return strings.TrimSpace(positional[0]), nil
	default:
		return "", fmt.Errorf("expected a single <problem> argument, got %q", positional)
	}
}

// ---------------------------------------------------------------------------
// Configuration
// ---------------------------------------------------------------------------

// findRepoRoot walks up from the working directory until it finds go.mod.
func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not locate repository root (no go.mod found)")
		}
		dir = parent
	}
}

// loadDotEnv reads KEY=VALUE pairs from path into the process environment.
// Existing environment variables take precedence over the file. Blank lines,
// "#" comments, an optional "export " prefix and surrounding quotes are handled.
func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return // .env is optional; plain environment variables also work
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if len(val) >= 2 && (val[0] == '"' || val[0] == '\'') && val[len(val)-1] == val[0] {
			val = val[1 : len(val)-1]
		}
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, val)
		}
	}
}

// ---------------------------------------------------------------------------
// LeetCode HTTP client
// ---------------------------------------------------------------------------

type leetClient struct {
	http    *http.Client
	session string
	csrf    string
}

func newLeetClient() *leetClient {
	return &leetClient{
		http:    &http.Client{Timeout: 30 * time.Second},
		session: os.Getenv("LEETCODE_SESSION"),
		csrf:    os.Getenv("LEETCODE_CSRFTOKEN"),
	}
}

// do performs a request (authenticated when cookies are available) and
// decodes the JSON response into out.
func (c *leetClient) do(method, url, referer string, body, out any) error {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, rdr)
	if err != nil {
		return err
	}
	if referer == "" {
		referer = baseURL + "/"
	}
	if c.session != "" && c.csrf != "" {
		req.Header.Set("Cookie", fmt.Sprintf("LEETCODE_SESSION=%s; csrftoken=%s", c.session, c.csrf))
		req.Header.Set("x-csrftoken", c.csrf)
	}
	req.Header.Set("Referer", referer)
	req.Header.Set("Origin", baseURL)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch {
	case resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden:
		return fmt.Errorf("HTTP %d from %s: cookies are likely expired or invalid — refresh LEETCODE_SESSION / LEETCODE_CSRFTOKEN in .env", resp.StatusCode, url)
	case resp.StatusCode == http.StatusTooManyRequests:
		return fmt.Errorf("HTTP 429 from %s: rate limited by LeetCode, wait a moment and retry", url)
	case resp.StatusCode >= 400:
		return fmt.Errorf("HTTP %d from %s: %s", resp.StatusCode, url, truncate(string(raw), 300))
	}

	if out == nil {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("unexpected response from %s (%v): %s", url, err, truncate(string(raw), 300))
	}
	return nil
}

func (c *leetClient) graphql(query string, vars map[string]any, out any) error {
	return c.do("POST", baseURL+"/graphql", "", map[string]any{"query": query, "variables": vars}, out)
}

// checkAuth verifies the cookies belong to a signed-in user.
func (c *leetClient) checkAuth() (string, error) {
	if c.session == "" || c.csrf == "" {
		return "", errors.New("LEETCODE_SESSION and LEETCODE_CSRFTOKEN must be set in .env for -test/-submit (see .env.example)")
	}
	var resp struct {
		Data struct {
			UserStatus struct {
				IsSignedIn bool   `json:"isSignedIn"`
				Username   string `json:"username"`
			} `json:"userStatus"`
		} `json:"data"`
	}
	if err := c.graphql(`query { userStatus { isSignedIn username } }`, nil, &resp); err != nil {
		return "", fmt.Errorf("checking authentication: %w", err)
	}
	if !resp.Data.UserStatus.IsSignedIn {
		return "", errors.New("not signed in: LEETCODE_SESSION cookie is missing, expired or invalid")
	}
	return resp.Data.UserStatus.Username, nil
}

// question holds everything needed to scaffold, run and submit a problem.
type question struct {
	ID           string // internal question_id used by the judge endpoints
	FrontendID   string
	Title        string
	Slug         string
	Difficulty   string
	Content      string // HTML problem statement
	GoSnippet    string // official Go starter code
	ExampleCases []string
	ParamCount   int // number of input lines per test case (0 if unknown)
}

func (c *leetClient) fetchQuestion(slug string) (*question, error) {
	const query = `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    title
    titleSlug
    difficulty
    content
    codeSnippets { langSlug code }
    exampleTestcaseList
    sampleTestCase
    metaData
  }
}`
	var resp struct {
		Data struct {
			Question *struct {
				QuestionID         string `json:"questionId"`
				QuestionFrontendID string `json:"questionFrontendId"`
				Title              string `json:"title"`
				TitleSlug          string `json:"titleSlug"`
				Difficulty         string `json:"difficulty"`
				Content            string `json:"content"`
				CodeSnippets       []struct {
					LangSlug string `json:"langSlug"`
					Code     string `json:"code"`
				} `json:"codeSnippets"`
				ExampleTestcaseList []string `json:"exampleTestcaseList"`
				SampleTestCase      string   `json:"sampleTestCase"`
				MetaData            string   `json:"metaData"`
			} `json:"question"`
		} `json:"data"`
	}
	if err := c.graphql(query, map[string]any{"titleSlug": slug}, &resp); err != nil {
		return nil, err
	}
	gq := resp.Data.Question
	if gq == nil || gq.QuestionID == "" {
		return nil, fmt.Errorf("problem %q not found on LeetCode", slug)
	}

	q := &question{
		ID:           gq.QuestionID,
		FrontendID:   gq.QuestionFrontendID,
		Title:        gq.Title,
		Slug:         gq.TitleSlug,
		Difficulty:   gq.Difficulty,
		Content:      gq.Content,
		ExampleCases: gq.ExampleTestcaseList,
	}
	for _, s := range gq.CodeSnippets {
		if s.LangSlug == "golang" {
			q.GoSnippet = s.Code
			break
		}
	}
	if len(q.ExampleCases) == 0 && gq.SampleTestCase != "" {
		q.ExampleCases = []string{gq.SampleTestCase}
	}

	var meta struct {
		Params []json.RawMessage `json:"params"`
	}
	if json.Unmarshal([]byte(gq.MetaData), &meta) == nil {
		q.ParamCount = len(meta.Params)
	}
	return q, nil
}

// checkResult models the payload of /submissions/detail/<id>/check/ for both
// "Run Code" (interpret) and "Submit" requests.
type checkResult struct {
	State              string          `json:"state"`
	StatusCode         int             `json:"status_code"`
	StatusMsg          string          `json:"status_msg"`
	RunSuccess         bool            `json:"run_success"`
	CorrectAnswer      *bool           `json:"correct_answer"`
	CodeAnswer         []string        `json:"code_answer"`
	ExpectedCodeAnswer []string        `json:"expected_code_answer"`
	StdOutputList      []string        `json:"std_output_list"`
	CompileError       string          `json:"compile_error"`
	FullCompileError   string          `json:"full_compile_error"`
	RuntimeError       string          `json:"runtime_error"`
	FullRuntimeError   string          `json:"full_runtime_error"`
	StatusRuntime      string          `json:"status_runtime"`
	StatusMemory       string          `json:"status_memory"`
	RuntimePercentile  *float64        `json:"runtime_percentile"`
	MemoryPercentile   *float64        `json:"memory_percentile"`
	TotalCorrect       *int            `json:"total_correct"`
	TotalTestcases     *int            `json:"total_testcases"`
	LastTestcase       string          `json:"last_testcase"`
	ExpectedOutput     string          `json:"expected_output"`
	CodeOutput         json.RawMessage `json:"code_output"`
	StdOutput          json.RawMessage `json:"std_output"`
}

// poll repeatedly queries the check endpoint until the judge finishes.
func (c *leetClient) poll(id string, q *question) (*checkResult, error) {
	url := fmt.Sprintf("%s/submissions/detail/%s/check/", baseURL, id)
	referer := fmt.Sprintf("%s/problems/%s/", baseURL, q.Slug)
	deadline := time.Now().Add(pollTimeout)

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for i := 0; ; i++ {
		var res checkResult
		if err := c.do("GET", url, referer, nil, &res); err != nil {
			fmt.Print("\r\033[K")
			return nil, err
		}
		if res.State == "SUCCESS" {
			fmt.Print("\r\033[K")
			return &res, nil
		}
		if time.Now().After(deadline) {
			fmt.Print("\r\033[K")
			return nil, fmt.Errorf("timed out waiting for judge (last state: %s)", res.State)
		}
		state := strings.ToLower(res.State)
		if state == "" {
			state = "queued"
		}
		fmt.Printf("\r%s judging… %s", cyan(frames[i%len(frames)]), dim(state))
		time.Sleep(pollInterval)
	}
}

// ---------------------------------------------------------------------------
// Target resolution
// ---------------------------------------------------------------------------

// problemRef identifies a problem and, when given as a path, its local files.
type problemRef struct {
	Num  int
	Slug string
	Dir  string // set when the target was a local path
	File string // set when the target was a local file
}

var dirNumRe = regexp.MustCompile(`^(\d{4})_`)

// resolveTarget maps a problem number, slug, directory or file path to the
// problem's number and LeetCode slug. Local folder names do not always match
// the slug (e.g. 0015_threesum vs "3sum"), so local lookups use the number.
func resolveTarget(target string) (*problemRef, error) {
	ref := &problemRef{}

	if info, err := os.Stat(target); err == nil && !isAllDigits(target) {
		abs, _ := filepath.Abs(target)
		if info.IsDir() {
			ref.Dir = abs
		} else {
			ref.File, ref.Dir = abs, filepath.Dir(abs)
		}
		m := dirNumRe.FindStringSubmatch(filepath.Base(ref.Dir))
		if m == nil {
			return nil, fmt.Errorf("cannot infer problem number from directory %q (expected NNNN_name)", ref.Dir)
		}
		ref.Num, _ = strconv.Atoi(m[1])
	} else if n, err := strconv.Atoi(target); err == nil {
		ref.Num = n
	} else {
		ref.Slug = strings.ReplaceAll(strings.ToLower(target), "_", "-")
		n, err := frontendIDForSlug(ref.Slug)
		if err != nil {
			return nil, err
		}
		ref.Num = n
	}

	if ref.Slug == "" {
		s, err := slugForNumber(ref.Num)
		if err != nil {
			return nil, err
		}
		ref.Slug = s
	}
	return ref, nil
}

func isAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// groupDirName returns the bucket folder for a problem number, e.g. 0010-0019.
func groupDirName(num int) string {
	if num >= 1 && num <= 9 {
		return "0001-0009"
	}
	start := (num / 10) * 10
	return fmt.Sprintf("%04d-%04d", start, start+9)
}

// findProblemDir locates <root>/<group>/<NNNN>_* for the given problem number.
func findProblemDir(root string, num int) (string, error) {
	matches, _ := filepath.Glob(filepath.Join(root, "*", fmt.Sprintf("%04d_*", num)))
	for _, m := range matches {
		if info, err := os.Stat(m); err == nil && info.IsDir() {
			return m, nil
		}
	}
	return "", fmt.Errorf("no local folder found for problem #%d — scaffold it first with: go run ./scripts/leetcode.go -fetch %d", num, num)
}

// pickSolutionFile returns the single non-test .go file in dir, preferring the
// one named after the folder when several exist.
func pickSolutionFile(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var candidates []string
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() && strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
			candidates = append(candidates, name)
		}
	}
	switch len(candidates) {
	case 0:
		return "", fmt.Errorf("no solution .go file found in %s", dir)
	case 1:
		return filepath.Join(dir, candidates[0]), nil
	}
	preferred := dirNumRe.ReplaceAllString(filepath.Base(dir), "") + ".go"
	if slices.Contains(candidates, preferred) {
		return filepath.Join(dir, preferred), nil
	}
	return "", fmt.Errorf("multiple .go files in %s (%s); pass the file path explicitly", dir, strings.Join(candidates, ", "))
}

// slugForNumber resolves a frontend problem number using the public problem list.
func slugForNumber(num int) (string, error) {
	req, err := http.NewRequest("GET", baseURL+"/api/problems/all/", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		StatStatusPairs []struct {
			Stat struct {
				FrontendQuestionID any    `json:"frontend_question_id"`
				Slug               string `json:"question__title_slug"`
			} `json:"stat"`
		} `json:"stat_status_pairs"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("decoding problem list: %w", err)
	}
	for _, p := range data.StatStatusPairs {
		var id int
		switch v := p.Stat.FrontendQuestionID.(type) {
		case float64:
			id = int(v)
		case string:
			id, _ = strconv.Atoi(v)
		}
		if id == num {
			return p.Stat.Slug, nil
		}
	}
	return "", fmt.Errorf("problem #%d not found on LeetCode", num)
}

// frontendIDForSlug resolves a slug to its frontend problem number.
func frontendIDForSlug(slug string) (int, error) {
	body, _ := json.Marshal(map[string]any{
		"query":     `query q($titleSlug: String!) { question(titleSlug: $titleSlug) { questionFrontendId } }`,
		"variables": map[string]string{"titleSlug": slug},
	})
	req, err := http.NewRequest("POST", baseURL+"/graphql", bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var out struct {
		Data struct {
			Question *struct {
				QuestionFrontendID string `json:"questionFrontendId"`
			} `json:"question"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, err
	}
	if out.Data.Question == nil {
		return 0, fmt.Errorf("problem %q not found on LeetCode", slug)
	}
	return strconv.Atoi(out.Data.Question.QuestionFrontendID)
}

// ---------------------------------------------------------------------------
// Fetch (scaffold)
// ---------------------------------------------------------------------------

// scaffold creates <group>/<NNNN>_<slug>/ with README.md, a starter solution
// and a test file. Existing non-empty files are never overwritten (except the
// README, which is regenerated). If a folder for the problem number already
// exists (possibly renamed, e.g. 0015_threesum), it is reused.
func scaffold(root string, q *question) error {
	num, _ := strconv.Atoi(q.FrontendID)
	folderSnake := strings.ReplaceAll(q.Slug, "-", "_")

	targetDir, err := findProblemDir(root, num)
	if err != nil {
		targetDir = filepath.Join(root, groupDirName(num), fmt.Sprintf("%04d_%s", num, folderSnake))
	} else {
		folderSnake = dirNumRe.ReplaceAllString(filepath.Base(targetDir), "")
	}
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return fmt.Errorf("creating %s: %w", targetDir, err)
	}
	rel := func(p string) string { r, _ := filepath.Rel(root, p); return filepath.ToSlash(r) }

	readmePath := filepath.Join(targetDir, "README.md")
	if err := os.WriteFile(readmePath, []byte(buildReadme(q)), 0o644); err != nil {
		return fmt.Errorf("writing README.md: %w", err)
	}
	fmt.Printf("%s %s\n", green("✓ wrote  "), rel(readmePath))

	pkgName := packageName(folderSnake)

	goPath := filepath.Join(targetDir, folderSnake+".go")
	if shouldWriteFile(goPath) {
		starter := fmt.Sprintf("package %s\n\n%s\n\n%s\n", pkgName, buildCommentBlock(q), q.GoSnippet)
		if err := os.WriteFile(goPath, []byte(starter), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", goPath, err)
		}
		fmt.Printf("%s %s\n", green("✓ created"), rel(goPath))
	} else {
		fmt.Printf("%s %s %s\n", dim("• skipped"), rel(goPath), dim("(already has content)"))
	}

	testPath := filepath.Join(targetDir, folderSnake+"_test.go")
	if shouldWriteFile(testPath) {
		starter := fmt.Sprintf("package %s\n\nimport (\n\t\"testing\"\n)\n\nfunc TestSolution(t *testing.T) {\n\t// TODO: add tests\n}\n", pkgName)
		if err := os.WriteFile(testPath, []byte(starter), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", testPath, err)
		}
		fmt.Printf("%s %s\n", green("✓ created"), rel(testPath))
	} else {
		fmt.Printf("%s %s %s\n", dim("• skipped"), rel(testPath), dim("(already has content)"))
	}
	return nil
}

var digitWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// packageName turns a snake_case folder name into a valid Go package name.
// Go identifiers cannot start with a digit, so leading digits are spelled out
// (e.g. "3sum_closest" -> "threesumclosest").
func packageName(folderSnake string) string {
	name := strings.ReplaceAll(folderSnake, "_", "")
	var prefix strings.Builder
	i := 0
	for i < len(name) && name[i] >= '0' && name[i] <= '9' {
		prefix.WriteString(digitWords[name[i]-'0'])
		i++
	}
	return prefix.String() + name[i:]
}

// shouldWriteFile reports whether path does not exist or is empty, so that
// existing solutions are never overwritten accidentally.
func shouldWriteFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	return err == nil && info.Size() == 0
}

// buildReadme formats the problem into a Markdown README structure.
func buildReadme(q *question) string {
	return fmt.Sprintf("# %s. %s\n\n**Difficulty:** %s\n\n%s\n", q.FrontendID, q.Title, q.Difficulty, htmlToMarkdown(q.Content))
}

// buildCommentBlock formats the problem description using "//" line comments
// for every line (including empty lines), adhering to Go documentation conventions.
func buildCommentBlock(q *question) string {
	fullText := fmt.Sprintf("%s. %s\n\n%s", q.FrontendID, q.Title, htmlToMarkdown(q.Content))
	lines := strings.Split(fullText, "\n")
	commented := make([]string, 0, len(lines))
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
// converting tags (<pre>, <code>, <strong>, <sup>, <li>, <p>) into GitHub Flavored Markdown.
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

	// Exponents: 10<sup>4</sup> -> 10^4 (otherwise it would render as "104")
	s = regexp.MustCompile(`<sup>(.*?)</sup>`).ReplaceAllString(s, "^$1")

	// Common inline formatting
	s = regexp.MustCompile(`<strong>(.*?)</strong>`).ReplaceAllString(s, "**$1**")
	s = regexp.MustCompile(`<b>(.*?)</b>`).ReplaceAllString(s, "**$1**")
	s = regexp.MustCompile(`<em>(.*?)</em>`).ReplaceAllString(s, "*$1*")
	s = regexp.MustCompile(`<i>(.*?)</i>`).ReplaceAllString(s, "*$1*")
	s = regexp.MustCompile(`<code>(.*?)</code>`).ReplaceAllString(s, "`$1`")

	// Paragraphs & line breaks
	for _, tag := range []string{"<p>", "</p>", "<br />", "<br/>", "<br>"} {
		s = strings.ReplaceAll(s, tag, "\n")
	}

	// Strip non-breaking spaces and HTML noise
	s = strings.ReplaceAll(s, "\u00a0", "")
	s = strings.ReplaceAll(s, "&nbsp;", "")

	// Unordered list items
	s = regexp.MustCompile(`<li>\s*(.*?)\s*</li>`).ReplaceAllString(s, "- $1\n")
	for _, tag := range []string{"<ul>", "</ul>", "<ol>", "</ol>"} {
		s = strings.ReplaceAll(s, tag, "\n")
	}

	// Strip any remaining HTML tags
	s = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(s, "")

	// Decode leftover HTML entities (&lt;, &gt;, etc.)
	s = html.UnescapeString(s)

	// Normalize list item indents
	s = regexp.MustCompile(`(?m)^\s+-\s+`).ReplaceAllString(s, "- ")

	// Clean up excessive blank lines
	s = regexp.MustCompile(`\n{3,}`).ReplaceAllString(s, "\n\n")

	return strings.TrimSpace(s)
}

// ---------------------------------------------------------------------------
// Code preprocessing
// ---------------------------------------------------------------------------

// prepareCode reads a local solution and converts it into the form LeetCode's
// Go judge expects: no package clause and no redeclaration of predefined
// helper types (ListNode, TreeNode, Node). Comments and imports are preserved.
func prepareCode(path string) (string, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("parse error (fix it locally first): %w", err)
	}

	type span struct{ start, end int }
	off := func(p token.Pos) int { return fset.Position(p).Offset }

	cuts := []span{{off(f.Package), off(f.Name.End())}}

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE || len(gd.Specs) == 0 {
			continue
		}
		allPredefined := true
		for _, s := range gd.Specs {
			if ts, ok := s.(*ast.TypeSpec); !ok || !predefinedTypes[ts.Name.Name] {
				allPredefined = false
				break
			}
		}
		if !allPredefined {
			continue
		}
		start := off(gd.Pos())
		if gd.Doc != nil {
			start = off(gd.Doc.Pos())
		}
		cuts = append(cuts, span{start, off(gd.End())})
	}

	slices.SortFunc(cuts, func(a, b span) int { return b.start - a.start })
	out := src
	for _, c := range cuts {
		out = append(out[:c.start:c.start], out[c.end:]...)
	}

	code := regexp.MustCompile(`\n{3,}`).ReplaceAllString(string(out), "\n\n")
	return strings.TrimSpace(code) + "\n", nil
}

// ---------------------------------------------------------------------------
// Test ("Run Code")
// ---------------------------------------------------------------------------

// buildTestInput returns the data_input payload and the individual cases used
// for per-case reporting.
func buildTestInput(q *question, input, inputFile string) (string, []string, error) {
	var raw string
	switch {
	case inputFile != "":
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return "", nil, err
		}
		raw = string(b)
	case input != "":
		raw = strings.ReplaceAll(input, `\n`, "\n")
	default:
		if len(q.ExampleCases) == 0 {
			return "", nil, errors.New("no example test cases available; supply -input or -input-file")
		}
		return strings.Join(q.ExampleCases, "\n"), q.ExampleCases, nil
	}

	var lines []string
	for _, l := range strings.Split(strings.ReplaceAll(raw, "\r\n", "\n"), "\n") {
		if strings.TrimSpace(l) != "" {
			lines = append(lines, strings.TrimSpace(l))
		}
	}
	if len(lines) == 0 {
		return "", nil, errors.New("custom input is empty")
	}

	per := q.ParamCount
	if per <= 0 || len(lines)%per != 0 {
		return strings.Join(lines, "\n"), []string{strings.Join(lines, "\n")}, nil
	}
	var cases []string
	for i := 0; i < len(lines); i += per {
		cases = append(cases, strings.Join(lines[i:i+per], "\n"))
	}
	return strings.Join(lines, "\n"), cases, nil
}

// runTest runs the code against the given cases and reports whether all passed.
func runTest(c *leetClient, q *question, code, dataInput string, cases []string) bool {
	var resp struct {
		InterpretID string `json:"interpret_id"`
	}
	payload := map[string]any{
		"question_id": q.ID,
		"data_input":  dataInput,
		"lang":        "golang",
		"typed_code":  code,
	}
	referer := fmt.Sprintf("%s/problems/%s/", baseURL, q.Slug)
	if err := c.do("POST", referer+"interpret_solution/", referer, payload, &resp); err != nil {
		fatalf("starting test run: %v", err)
	}
	if resp.InterpretID == "" {
		fatalf("LeetCode did not return an interpret_id")
	}

	res, err := c.poll(resp.InterpretID, q)
	if err != nil {
		fatalf("%v", err)
	}

	if printJudgeError(res) {
		return false
	}

	allPassed := true
	for i, tc := range cases {
		got, want := at(res.CodeAnswer, i), at(res.ExpectedCodeAnswer, i)
		ok := got == want
		allPassed = allPassed && ok

		mark := green("✓ Case " + strconv.Itoa(i+1))
		if !ok {
			mark = red("✗ Case " + strconv.Itoa(i+1))
		}
		fmt.Println(bold(mark))
		fmt.Printf("  %s %s\n", dim("Input:   "), indent(tc, "           "))
		fmt.Printf("  %s %s\n", dim("Output:  "), colorIf(ok, got))
		fmt.Printf("  %s %s\n", dim("Expected:"), want)
		if s := strings.TrimSpace(at(res.StdOutputList, i)); s != "" {
			fmt.Printf("  %s %s\n", dim("Stdout:  "), indent(s, "           "))
		}
		fmt.Println()
	}

	if allPassed {
		fmt.Printf("%s  %s\n", green(bold("✓ All sample cases passed")), dim("runtime "+res.StatusRuntime))
	} else {
		fmt.Println(red(bold("✗ Wrong Answer on sample cases")))
	}
	return allPassed
}

// ---------------------------------------------------------------------------
// Submit
// ---------------------------------------------------------------------------

// runSubmit submits the code and reports whether it was accepted.
func runSubmit(c *leetClient, q *question, code string) bool {
	var resp struct {
		SubmissionID json.Number `json:"submission_id"`
	}
	payload := map[string]any{
		"question_id": q.ID,
		"lang":        "golang",
		"typed_code":  code,
	}
	referer := fmt.Sprintf("%s/problems/%s/", baseURL, q.Slug)
	if err := c.do("POST", referer+"submit/", referer, payload, &resp); err != nil {
		fatalf("submitting: %v", err)
	}
	if resp.SubmissionID == "" {
		fatalf("LeetCode did not return a submission_id")
	}

	res, err := c.poll(resp.SubmissionID.String(), q)
	if err != nil {
		fatalf("%v", err)
	}

	link := fmt.Sprintf("%s/submissions/detail/%s/", baseURL, resp.SubmissionID)

	if printJudgeError(res) {
		printPassed(res)
		fmt.Println(dim("  " + link))
		return false
	}

	if res.StatusCode == 10 {
		fmt.Println(green(bold("✓ Accepted")))
		printPassed(res)
		fmt.Printf("  %s %-10s %s\n", dim("Runtime:"), res.StatusRuntime, percentile(res.RuntimePercentile))
		fmt.Printf("  %s  %-10s %s\n", dim("Memory:"), res.StatusMemory, percentile(res.MemoryPercentile))
		fmt.Println(dim("  " + link))
		return true
	}

	// Wrong Answer, TLE, MLE, OLE, etc.
	fmt.Println(red(bold("✗ " + res.StatusMsg)))
	printPassed(res)
	if res.LastTestcase != "" {
		fmt.Printf("  %s %s\n", dim("Input:   "), indent(res.LastTestcase, "           "))
	}
	if out := rawToString(res.CodeOutput); out != "" {
		fmt.Printf("  %s %s\n", dim("Output:  "), red(out))
	}
	if res.ExpectedOutput != "" {
		fmt.Printf("  %s %s\n", dim("Expected:"), res.ExpectedOutput)
	}
	if out := rawToString(res.StdOutput); out != "" {
		fmt.Printf("  %s %s\n", dim("Stdout:  "), indent(out, "           "))
	}
	fmt.Println(dim("  " + link))
	return false
}

// ---------------------------------------------------------------------------
// Output helpers
// ---------------------------------------------------------------------------

// printJudgeError prints compile/runtime errors and reports whether one occurred.
func printJudgeError(res *checkResult) bool {
	switch {
	case res.CompileError != "" || res.StatusCode == 20:
		fmt.Println(yellow(bold("! Compile Error")))
		fmt.Println(indent(firstNonEmpty(res.FullCompileError, res.CompileError), "  "))
		return true
	case res.RuntimeError != "" || res.StatusCode == 15:
		fmt.Println(red(bold("✗ Runtime Error")))
		fmt.Println(indent(firstNonEmpty(res.FullRuntimeError, res.RuntimeError), "  "))
		if res.LastTestcase != "" {
			fmt.Printf("  %s %s\n", dim("Last input:"), indent(res.LastTestcase, "             "))
		}
		return true
	}
	return false
}

func printPassed(res *checkResult) {
	if res.TotalCorrect != nil && res.TotalTestcases != nil {
		fmt.Printf("  %s %d / %d test cases passed\n", dim("Passed: "), *res.TotalCorrect, *res.TotalTestcases)
	}
}

func percentile(p *float64) string {
	if p == nil {
		return ""
	}
	return cyan(fmt.Sprintf("beats %.2f%%", *p))
}

func colorIf(ok bool, s string) string {
	if ok {
		return green(s)
	}
	return red(s)
}

func at(s []string, i int) string {
	if i < len(s) {
		return s[i]
	}
	return ""
}

func indent(s, prefix string) string {
	return strings.ReplaceAll(strings.TrimRight(s, "\n"), "\n", "\n"+prefix)
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// rawToString renders a JSON field that may be a string or an array of strings.
func rawToString(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return s
	}
	var list []string
	if json.Unmarshal(raw, &list) == nil {
		return strings.Join(list, "\n")
	}
	return string(raw)
}

func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}
