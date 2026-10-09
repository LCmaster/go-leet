# LeetCode Go Solutions

This repository contains my solutions to LeetCode problems written in Go.

## Structure

The repository is organized by problem number and name. Each problem directory contains:
- The Go solution file (e.g., `two_sum.go`)
- A Go test file for table-driven testing (e.g., `two_sum_test.go`)

## How to Run Tests

To run the tests for a specific problem, navigate to its directory and run:

```bash
cd 0001_two_sum
go test -v
```

To run all tests in the repository:

```bash
go test ./...
```

## Problems Solved

| # | Title | Solution | Difficulty |
|---|-------|----------|------------|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [two_sum.go](./0001_two_sum/two_sum.go) | Easy |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | [add_two_numbers.go](./0002_add_two_numbers/add_two_numbers.go) | Medium |
| 3 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [longest_substring.go](./0003_longest_substring/longest_substring.go) | Medium |
