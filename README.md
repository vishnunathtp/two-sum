# Two Sum

Fast $O(N)$ implementation in Go solving the classic Two Sum problem using a hash map.

## Problem Description
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

### Example
- Input: `nums = [2, 7, 11, 15], target = 9`
- Output: `[0, 1]`

## Approach & Complexity
Use a hash map to store each visited element and its index. For each number, check if the complement (`target - num`) already exists in the map.

- **Time Complexity:** $O(N)$ single pass.
- **Space Complexity:** $O(N)$ for hash map storage.

## How to Run & Test
```bash
go run twosum.go
go test -v ./...
```
