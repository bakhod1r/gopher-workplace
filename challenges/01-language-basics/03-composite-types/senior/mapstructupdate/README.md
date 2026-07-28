# Update Struct in Map

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`m[key]` returns a **copy** of the struct. Incrementing `s.Hits` changes the copy;
without writing it back, the map keeps the old value. (Go even forbids
`m[key].Hits++` directly.)

## Task

Fix the line between the markers in
[mapstructupdate.go](mapstructupdate.go) to store the updated struct.

## Examples

```go
Record(m,"a"); Record(m,"a") // m["a"].Hits == 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map value copy** | `m[key]` is not addressable. |
| 2 | **Read-modify-write** | Copy, mutate, assign back. |
| 3 | **No `m[k].F++`** | Compile error; that's why. |

## Hint

`m[key] = s`.

## Validate

```bash
make verify
```
