# LeetCode Go Solutions

This repository contains my solutions to LeetCode problems written in Go.

## Structure

The repository is organized by problem number and name. Each problem directory contains:
- The Go solution file (e.g., `two_sum.go`)
- A Go test file for table-driven testing (e.g., `two_sum_test.go`)

## How to Run Tests

To run the tests for a specific problem, navigate to its directory and run:

```bash
cd 0001-0009/0001_two_sum
go test -v
```

To run all tests in the repository:

```bash
go test ./...
```

## How to Fetch a New Problem

You can download problem descriptions, examples, constraints, and starter code automatically:

```bash
go run ./scripts/fetch.go <problem_number_or_slug>
```

Examples:
```bash
go run ./scripts/fetch.go 5
go run ./scripts/fetch.go longest-palindromic-substring
```

## Problems Solved

| # | Title | Solution | Difficulty |
|---|-------|----------|------------|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [two_sum.go](./0001-0009/0001_two_sum/two_sum.go) | Easy |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | [add_two_numbers.go](./0001-0009/0002_add_two_numbers/add_two_numbers.go) | Medium |
| 3 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [longest_substring.go](./0001-0009/0003_longest_substring/longest_substring.go) | Medium |
| 4 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | [mediam_of_two_sorted_arrays.go](./0001-0009/0004_mediam_of_two_sorted_arrays/mediam_of_two_sorted_arrays.go) | Hard |
| 5 | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | [longest_palindromic_substring.go](./0001-0009/0005_longest_palindromic_substring/longest_palindromic_substring.go) | Medium |
| 6 | [Zigzag Conversion](https://leetcode.com/problems/zigzag-conversion/) | [zigzag_conversion.go](./0001-0009/0006_zigzag_conversion/zigzag_conversion.go) | Medium |
| 7 | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | [reverse_integer.go](./0001-0009/0007_reverse_integer/reverse_integer.go) | Medium |
| 8 | [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) | [string_to_integer_atoi.go](./0001-0009/0008_string_to_integer_atoi/string_to_integer_atoi.go) | Medium |
| 9 | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | [palindrome_number.go](./0001-0009/0009_palindrome_number/palindrome_number.go) | Easy |
