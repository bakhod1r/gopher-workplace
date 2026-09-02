# Equal By Predicate

**Level:** middle  
**Topic:** 03-generics

## Context

A test compares parsed records against expected fixtures whose types differ, so `==` is not available.

## Task

Implement the stub(s) in [equalbygen.go](equalbygen.go):

1. Implement `EqualBy`, reporting whether the slices match pairwise under `eq`.
2. Different lengths are never equal; two empty slices are.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  EqualBy([]int{1}, []string{"1"}, matches)
Output: true
```

**Example 2:**

```
Input:  EqualBy([]int{1}, []string{"2"}, matches)
Output: false
```

**Example 3:**

```
Input:  EqualBy(nil, []string{}, matches)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cross-type comparison** | Two element type parameters let the two slices differ entirely. |
| 2 | **Custom equality** | The predicate replaces `==`, so neither element type needs `comparable`. |
| 3 | **Length guard first** | It makes the index loop safe and short-circuits the common case. |

## Hint

Two type parameters; `eq` bridges them.

## Validate

```bash
make verify
```
