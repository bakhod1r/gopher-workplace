# Double indirection

## Intuition

A `**int` lets a callee reseat the caller's pointer variable; `*pp = q` stores a new address into it.

## Approach

1. `pp` is a `**int`: it points at the caller's pointer variable.
2. `*pp = q` writes a new address into that pointer.
3. This changes where `p` points, not the int it pointed at.

## Solution

```go
func Reseat(pp **int, q *int) {
	*pp = q
}
```

## Walkthrough

`Reseat(&p, &b)`: `*pp` refers to `p`; assigning `q` (which is `&b`) makes `p` point at `b`.

## Pitfalls

- Reassigning a plain `*int` parameter only changes the local copy.
- `**int` reaches the caller's pointer.
