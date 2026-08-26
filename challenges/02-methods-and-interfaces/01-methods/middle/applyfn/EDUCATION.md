# Method Values as Callbacks

## Intuition

Because `tr.Transform` is already a `func(int) int`, you can pass it to any
function that accepts that signature. No wrapper needed — Go's method values
make this seamless.

## Approach

1. Create a result slice.
2. Range over `nums`, apply `fn`, append.
3. Return result.

## Solution

```go
func ApplyAll(fn func(int) int, nums []int) []int {
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		result = append(result, fn(n))
	}
	return result
}
```

## Walkthrough

For `fn = Transformer{2}.Transform`, `nums = [1, 2, 3]`:
- fn(1) = 2, fn(2) = 4, fn(3) = 6.
- result = [2, 4, 6].

## Pitfalls

- Returning `nil` instead of `[]int{}` for empty input — `reflect.DeepEqual`
  distinguishes nil from empty.
- Forgetting that method values capture the receiver — the `Factor` is baked in.
