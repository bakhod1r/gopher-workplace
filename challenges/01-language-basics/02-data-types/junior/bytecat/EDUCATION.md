# String ↔ []byte conversion

## The idea

Strings and byte slices convert with an ordinary conversion:

```go
[]byte("Go")          // []byte{71, 111}
string([]byte{71, 111}) // "Go"
```

A string is an immutable sequence of bytes; a `[]byte` is a mutable one. Each
conversion **copies**, so mutating the slice never changes the source string.

## Why it matters

I/O, hashing, and encoding APIs work in `[]byte`; text APIs work in `string`.
Converting between them is constant, everyday plumbing — and the copy is what
keeps string immutability guarantees intact.

## Watch out

- Each conversion allocates and copies (O(n)); avoid it in hot loops.
- Converting a `[]byte` with invalid UTF-8 to a string keeps the bytes as-is;
  ranging over it later yields `utf8.RuneError` for bad sequences.
- `string(65)` is `"A"` (rune conversion), *not* `"65"` — a common surprise;
  `string([]byte{65})` is also `"A"`.

## Try it yourself

```go
string([]byte{72, 105}) // "Hi"
[]byte("Hi")            // [72 105]
string(rune(65))        // "A"
```
