# Formatter

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A logging front-end renders one record either as plain text or as a key=value line.

## Task

Implement the stub(s) in [formatter.go](formatter.go):

1. Implement `Format` on `Plain` — return the message unchanged.
2. Implement `Format` on `KeyValue` — return `"msg=<message>"`.
3. Implement `Render`, which formats the message with the given formatter.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Plain{}.Format("hi")
Output: "hi"
```

**Example 2:**

```
Input:  KeyValue{}.Format("hi")
Output: "msg=hi"
```

**Example 3:**

```
Input:  Render(KeyValue{}, "boom")
Output: "msg=boom"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strategy via interface** | Swapping the implementation swaps the behaviour with no branching. |
| 2 | **Empty structs as behaviour holders** | `Plain{}` has no state — only a method. |
| 3 | **String concatenation** | Reused: `+` or `fmt.Sprintf`. |

## Hint

`Render` must not switch on the type — call `f.Format(msg)`.

## Validate

```bash
make verify
```
