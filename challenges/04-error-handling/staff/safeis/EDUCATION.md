# Comparison That Cannot Panic

## Intuition

`errors.Is` performs `==` on interface values, which panics when the dynamic type is uncomparable. Framework code that accepts arbitrary errors has to assume the worst.

## Approach

1. Return false when either argument is nil.
2. Defer a recovery that leaves `match` false.
3. Return `errors.Is(err, target)`.

## Solution

```go
if err == nil || target == nil {
	return false
}
defer func() {
	if recover() != nil {
		match = false
	}
}()
return errors.Is(err, target)
```

## Walkthrough

Comparing the slice-bearing struct against itself panics inside `errors.Is`; the deferred recovery converts that into a plain false.

## Pitfalls

- Using an unnamed result, so the recovered path cannot set the answer.
- Recovering around the whole request instead of the single comparison.
- Assuming every error type is comparable because most are.
