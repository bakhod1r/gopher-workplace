# Recursive Methods

## Intuition

Methods can call themselves recursively. For a tree, `t.Walk()` calls
`t.Left.Walk()` and `t.Right.Walk()`. The nil receiver is the base case.

## Approach

1. Base case: `t == nil` → return `[]int{}`.
2. Recurse left, append `t.Val`, recurse right.

## Solution

```go
func (t *Tree) Walk() []int {
	if t == nil {
		return []int{}
	}
	result := t.Left.Walk()
	result = append(result, t.Val)
	result = append(result, t.Right.Walk()...)
	return result
}
```

## Walkthrough

For `Tree{2, Tree{1}, Tree{3}}`:
- Left: `Tree{1}.Walk()` → [1].
- Root: append 2 → [1, 2].
- Right: `Tree{3}.Walk()` → [3].
- Result: [1, 2, 3].

## Pitfalls

- Forgetting nil check → panic on `t.Left` when `t` is nil.
- Returning `nil` instead of `[]int{}` — `reflect.DeepEqual` distinguishes them.
