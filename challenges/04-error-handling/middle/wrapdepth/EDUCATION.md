# Chain Length

## Intuition

Walking a chain and counting it are the same loop; only the bookkeeping differs. The nil at the end terminates the walk and is not itself a link.

## Approach

1. Start a counter at 0.
2. While `err != nil`, increment and unwrap.
3. Return the counter.

## Solution

```go
n := 0
for err != nil {
	n++
	err = errors.Unwrap(err)
}
return n
```

## Walkthrough

For two wraps the loop counts the outer error, the middle wrapper and `ErrBase`, then `Unwrap` returns nil and the loop stops at 3.

## Pitfalls

- Starting the loop after one unwrap, undercounting by one.
- Counting the terminating nil as a link.
- Assuming `fmt.Errorf` without `%w` produces a wrapper — it does not.
