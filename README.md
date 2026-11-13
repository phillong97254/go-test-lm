# Go Live Coding Assignment — Word Frequency Counter

**Time limit:** 5–10 minutes

---

## Overview

You will implement a simple **in-memory word frequency store** in Go.

The store tracks how many times each word has been seen, and can report the most frequent words.

---

## Your Tasks

Work inside `store.go`. Three functions need to be implemented — their signatures are already provided.

### 1. `Add(word string)`
Record one occurrence of the given word.
- Words should be case-insensitive (`"Go"` and `"go"` count as the same word).

### 2. `Count(word string) int`
Return how many times the given word has been added.
- Return `0` for words that have never been added.

### 3. `TopN(n int) []string`
Return the `n` most frequent words, sorted from highest to lowest count.
- If two words have the same count, either order between them is acceptable.
- If `n` is greater than the number of unique words, return all words.

---

## Files

| File | Description |
|------|-------------|
| `main.go` | Entry point — you can use this to manually test your implementation |
| `store.go` | **Your work goes here** — implement the three functions |
| `store_test.go` | Test suite — run this to verify your implementation |

---

## Running the Code

```bash
# Run the program
go run .

# Run the tests
go test ./...
```

---

## Rules

- Only edit `store.go`.
- Do not modify `store_test.go` or `main.go`.
- Standard library only — no third-party packages.
