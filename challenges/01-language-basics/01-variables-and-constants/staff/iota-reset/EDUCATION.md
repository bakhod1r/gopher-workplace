# iota and implicit repetition

## The idea

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

## Why it matters

The mental model "iota just counts, so the names stay in order" is wrong once an
explicit value appears. The break is subtle: `iota` itself keeps advancing, but
the *bound values* freeze at the last written expression.

## Watch out

- Implicit repetition copies the whole RHS expression, including a literal.
- If you must inject a specific value, re-establish the pattern on the next line
  (`= iota`) or the run stays frozen.
- Reordering members silently renumbers them — enums are position-based.

## Try it yourself

```go
const (
	A = iota // 0
	B        // 1
	C = 9    // 9
	D        // 9  (repeats "= 9")
	E = iota // 4  (back on the counter)
)
```
