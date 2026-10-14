package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepareCode(t *testing.T) {
	src := `package addtwonumbers

// Problem description.

import "fmt"

// ListNode is predefined on LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

type helper struct{}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	fmt.Println(helper{})
	return nil
}
`
	path := filepath.Join(t.TempDir(), "sol.go")
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := prepareCode(path)
	if err != nil {
		t.Fatal(err)
	}

	for _, banned := range []string{"package ", "type ListNode", "predefined on LeetCode"} {
		if strings.Contains(got, banned) {
			t.Errorf("output still contains %q:\n%s", banned, got)
		}
	}
	for _, kept := range []string{"// Problem description.", `import "fmt"`, "type helper struct{}", "func addTwoNumbers"} {
		if !strings.Contains(got, kept) {
			t.Errorf("output is missing %q:\n%s", kept, got)
		}
	}
}

func TestLoadDotEnv(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	content := "# comment\nexport GOLEET_A=\"quoted\"\nGOLEET_B = plain \n\nGOLEET_C='single'\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"GOLEET_A", "GOLEET_B", "GOLEET_C"} {
		os.Unsetenv(k)
		t.Cleanup(func() { os.Unsetenv(k) })
	}

	loadDotEnv(path)

	want := map[string]string{"GOLEET_A": "quoted", "GOLEET_B": "plain", "GOLEET_C": "single"}
	for k, v := range want {
		if got := os.Getenv(k); got != v {
			t.Errorf("%s = %q, want %q", k, got, v)
		}
	}
}

func TestPackageName(t *testing.T) {
	tests := map[string]string{
		"two_sum":      "twosum",
		"3sum":         "threesum",
		"3sum_closest": "threesumclosest",
		"4sum_ii":      "foursumii",
	}
	for in, want := range tests {
		if got := packageName(in); got != want {
			t.Errorf("packageName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestGroupDirName(t *testing.T) {
	tests := map[int]string{1: "0001-0009", 9: "0001-0009", 10: "0010-0019", 16: "0010-0019", 120: "0120-0129"}
	for in, want := range tests {
		if got := groupDirName(in); got != want {
			t.Errorf("groupDirName(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestHTMLToMarkdownSup(t *testing.T) {
	got := htmlToMarkdown("<code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code>")
	if want := "`-10^4 <= target <= 10^4`"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
