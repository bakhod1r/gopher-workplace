# iota and implicit repetition

## Intuition

`iota` increments on every `ConstSpec` line, but a **bare** line does not repeat
`iota` — it repeats the *previous line's expression* verbatim. Put an explicit
value in the middle and the bare lines below copy **that constant**, not the
counter:

```go
const (
	Info    Class = iota // 0
	Success              // 1
	Redirect = 7         // explicit; iota is now 2 but unused here
	ClientError          // repeats "= 7"  -> 7, NOT 3
	ServerError          // repeats "= 7"  -> 7
)
```

Keep every member on the `iota` expression to preserve the sequence:

```go
Redirect // bare -> repeats "iota" -> 2
```

## Approach

1. Assigning `Redirect = 7` overrides the running `iota` and the repeat below it.
2. Use `Redirect Class = iota` to keep the sequence 2, 3, 4.

## Solution

```go
type Class int

const (
	Info    Class = iota // 0
	Success              // 1
	Redirect Class = iota
	ClientError // should be 3
	ServerError // should be 4
)
```

## Walkthrough

The literal 7 makes ClientError repeat 7 too. `iota` restores Redirect=2, so ClientError=3, ServerError=4.

## Pitfalls

- Implicit repetition copies the whole RHS expression, including a literal.
- If you must inject a specific value, re-establish the pattern on the next line
  (`= iota`) or the run stays frozen.
- Reordering members silently renumbers them — enums are position-based.
