# Double-pointer reassignment

## Intuition

`pp` is a copy of the address of the caller's pointer; only `*pp = q` writes through to reseat it.

## Approach

1. `pp` is a `**int`. The bug writes `pp = &q`, rebinding the local parameter only.
2. `*pp = q` writes through to the caller's pointer.

## Solution

```go
func Reseat(pp **int, q *int) {
	*pp = q
}
```

## Walkthrough

`pp = &q` changes where `pp` itself points inside the function; the caller's `p` is untouched. `*pp = q` stores `q` into `p`.

## Pitfalls

- `pp = &q` changes a local copy.
- `*pp = q` reseats the caller's pointer.
