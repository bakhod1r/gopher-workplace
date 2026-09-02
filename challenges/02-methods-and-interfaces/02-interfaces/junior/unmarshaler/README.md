# Unmarshaler

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An import job parses wire text back into domain values and reports malformed input.

## Task

Implement the stub(s) in [unmarshaler.go](unmarshaler.go):

1. Implement `Unmarshal` on `*Pair` — parse `"<key>=<value>"`, returning an error when the `=` is missing.
2. Implement `UnmarshalAll`, which parses every line and stops at the first error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  p := &Pair{}; p.Unmarshal("a=1")
Output: nil, p == Pair{Key: "a", Value: "1"}
```

**Example 2:**

```
Input:  p.Unmarshal("nope")
Output: error "bad pair"
```

**Example 3:**

```
Input:  UnmarshalAll([]string{"a=1", "b=2"})
Output: 2 pairs, nil error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Parsing behind an interface** | The mirror image of `Marshal`: text in, state out. |
| 2 | **strings.Cut** | Reused from standard library basics: split on the first separator. |
| 3 | **Error propagation** | Reused: return the first failure to the caller. |

## Hint

`strings.Cut(s, "=")` returns `before, after, found`.

## Validate

```bash
make verify
```
