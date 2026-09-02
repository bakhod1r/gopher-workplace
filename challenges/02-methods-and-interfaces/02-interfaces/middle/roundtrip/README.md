# Round Trip

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A sync tool serialises records, ships them, and parses them back, checking that nothing was lost.

## Task

Implement the stub(s) in [roundtrip.go](roundtrip.go):

1. Implement `Marshal`/`Unmarshal` on `*Record` using the `"<id>|<name>"` format.
2. Implement `RoundTrip`, which marshals a codec, unmarshals into a fresh one, and reports whether the text survived.
3. Return `ErrBadFormat` for input without a `|` or with a non-numeric id.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&Record{ID: 1, Name: "a"}).Marshal()
Output: "1|a"
```

**Example 2:**

```
Input:  r.Unmarshal("2|b")
Output: nil, r == Record{2, "b"}
```

**Example 3:**

```
Input:  RoundTrip(&Record{ID: 3, Name: "c"}, &Record{})
Output: true, nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two-way interface** | `Codec` combines `Marshaler` and `Unmarshaler`. |
| 2 | **strconv.Atoi errors** | Reused from standard library: parse failures are values. |
| 3 | **strings.Cut** | Reused: split on the first separator. |

## Hint

`RoundTrip` compares `src.Marshal()` with `dst.Marshal()` after the parse.

## Validate

```bash
make verify
```
