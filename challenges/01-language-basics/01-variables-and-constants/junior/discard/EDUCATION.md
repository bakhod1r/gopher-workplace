# The blank identifier

## The idea

`_` is a name you are allowed to assign to and never allowed to read. It exists
because Go requires you to account for every value — and sometimes the honest
answer is "I do not want this one".

```go
pages, _ := Split(10, 3)   // keep the count, drop the remainder
_, rest := Split(10, 3)    // the other way round
```

Assigning to `_` compiles to nothing. It is not a variable: it has no address,
no type, and cannot be read back.

## Why it is needed

Two Go rules make it necessary rather than merely convenient.

**A multi-value call must be received in full.** There is no "take the first"
shorthand:

```go
pages := Split(10, 3)      // compile error: 2 values, 1 variable
```

**Unused local variables are a compile error.** So you cannot dodge the first
rule by naming the value and ignoring it:

```go
pages, rest := Split(10, 3)
return pages               // compile error: rest declared and not used
```

`_` satisfies the first rule without triggering the second.

## Where else it shows up

```go
for _, v := range items { }        // index not needed
for i := range items { }           // value not needed (no _ required)

var _ io.Writer = (*MyType)(nil)   // compile-time interface check
import _ "net/http/pprof"          // import for side effects only
func (Handler) ServeHTTP(_ http.ResponseWriter, r *http.Request) { }
```

The last one is worth noting: an unused *parameter* is legal in Go, so naming it
`_` is a message to the reader, not a requirement.

## Not a licence to ignore errors

```go
data, _ := os.ReadFile(path)   // the error is real; you just hid it
```

This compiles and is almost always a bug — `data` is nil and nothing said why.
Discard a value when it is genuinely irrelevant, not when handling it is
inconvenient. Linters flag `_` on error returns for good reason.

## Watch out

- `_` cannot be read: `x := _` does not compile.
- `_ = someExpr` is a legal statement, occasionally used to silence "declared
  and not used" during debugging. Remove it before committing.
- Each `_` is independent; `_, _ := f()` is fine but pointless — with nothing
  new on the left, use `f()` alone if it is called for effect.
- `_` in a `range` is only needed for the *value*: `for i := range xs` already
  drops it.

## Try it yourself

```go
pages, _ := Split(10, 3)
_, rest := Split(10, 3)
fmt.Println(pages, rest)      // 3 1

for _, v := range []string{"a", "b"} {
	fmt.Println(v)            // a b — index discarded
}
```
