# LeetCode Go Solutions

[![Solved](https://img.shields.io/badge/Solved-19-2ea44f?style=flat-square)](#problems-solved)
[![Easy](https://img.shields.io/badge/Easy-4-00b8a3?style=flat-square)](#problems-solved)
[![Medium](https://img.shields.io/badge/Medium-13-ffc01e?style=flat-square)](#problems-solved)
[![Hard](https://img.shields.io/badge/Hard-2-ff375f?style=flat-square)](#problems-solved)

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

## LeetCode CLI

A single script handles fetching, testing and submitting problems:

```bash
go run ./scripts/leetcode.go [-fetch] [-test] [-submit] <problem>
```

`<problem>` can be a problem number (`16`), a slug (`3sum-closest`) or a path to the problem
folder/file. Flags can be combined; they run in the order fetch → test → submit, and submit is
skipped if the test fails.

### Fetch a new problem

Downloads the description, examples, constraints and Go starter code, and scaffolds the folder:

```bash
go run ./scripts/leetcode.go -fetch 16
```

### Configure cookies (required for `-test` / `-submit`)

Copy `.env.example` to `.env` (it is git-ignored) and fill in:

```env
LEETCODE_SESSION=...
LEETCODE_CSRFTOKEN=...
```

To get the values, log in to [leetcode.com](https://leetcode.com), open DevTools (F12) →
**Application** (Chrome) / **Storage** (Firefox) → **Cookies** → `https://leetcode.com`,
and copy the `LEETCODE_SESSION` and `csrftoken` cookies. They expire periodically; refresh them when the
script reports you are not signed in.

### Run against sample test cases ("Run Code")

```bash
go run ./scripts/leetcode.go -test 16
go run ./scripts/leetcode.go -test 16 -input "[-1,2,1,-4]\n1"
go run ./scripts/leetcode.go -test 16 -input-file cases.txt
```

### Submit

```bash
go run ./scripts/leetcode.go -submit 16
go run ./scripts/leetcode.go -test -submit 16   # submit only if the samples pass
```

Before uploading, the `package` clause and local declarations of LeetCode's predefined types
(`ListNode`, `TreeNode`, `Node`) are stripped automatically. Upon an `Accepted` submission, the
README.md table and progress badges are updated automatically.

### Progress & Stats Tracker

View a terminal breakdown of your repository progress and live LeetCode account statistics:

```bash
go run ./scripts/leetcode.go -stats
```

### Sync README.md

Rescan all local solution directories and regenerate the statistics badges and "Problems Solved" table:

```bash
go run ./scripts/leetcode.go -sync
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
| 10 | [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) | [regular_expression_matching.go](./0010-0019/0010_regular_expression_matching/regular_expression_matching.go) | Hard |
| 11 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | [container_with_most_water.go](./0010-0019/0011_container_with_most_water/container_with_most_water.go) | Medium |
| 12 | [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) | [integer_to_roman.go](./0010-0019/0012_integer_to_roman/integer_to_roman.go) | Medium |
| 13 | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | [roman_to_integer.go](./0010-0019/0013_roman_to_integer/roman_to_integer.go) | Easy |
| 14 | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | [longest_common_prefix.go](./0010-0019/0014_longest_common_prefix/longest_common_prefix.go) | Easy |
| 15 | [3Sum](https://leetcode.com/problems/3sum/) | [threesum.go](./0010-0019/0015_threesum/threesum.go) | Medium |
| 16 | [3Sum Closest](https://leetcode.com/problems/3sum-closest/) | [threesum_closest.go](./0010-0019/0016_threesum_closest/threesum_closest.go) | Medium |
| 17 | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [letter_combinations_of_a_phone_number.go](./0010-0019/0017_letter_combinations_of_a_phone_number/letter_combinations_of_a_phone_number.go) | Medium |
| 18 | [4Sum](https://leetcode.com/problems/4sum/) | [foursum.go](./0010-0019/0018_foursum/foursum.go) | Medium |
| 19 | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [remove_nth_node_from_end_of_list.go](./0010-0019/0019_remove_nth_node_from_end_of_list/remove_nth_node_from_end_of_list.go) | Medium |
