# Struct tags and JSON

## The idea

A struct tag is metadata read by libraries via reflection. `encoding/json` uses
the `json:"..."` tag as the exact output key:

```go
LastName string `json:"last_name"`
```

## Why it matters

Tags define your serialization contract. A typo (`lastName` vs `last_name`)
silently ships the wrong key — the code compiles and marshals fine, but consumers
can't find the field. Tag correctness is an API concern.

## Watch out

- The tag string is literal; casing and spelling matter.
- Only **exported** (capitalized) fields are marshaled.
- Options like `,omitempty` follow the name: `json:"x,omitempty"`.
