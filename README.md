# Find the Missing Number

Bitwise XOR implementation in Go to find the missing number in an array containing $n$ distinct numbers in range $[0, n]$.

## Problem Description
Given an array containing $n$ distinct numbers taken from $0, 1, 2, \dots, n$, find the one that is missing from the array.

### Example
- Input: `[3, 0, 1]` -> Output: `2`

## Approach & Complexity
Bitwise XOR property: $x \oplus x = 0$ and $x \oplus 0 = x$. XORing all array indices with array elements cancels all present numbers, leaving only the missing one.

- **Time Complexity:** $O(N)$
- **Space Complexity:** $O(1)$ (no extra space, avoids arithmetic overflow).

## How to Run & Test
```bash
go run missing.go
go test -v ./...
```
