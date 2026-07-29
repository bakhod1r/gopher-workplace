# Zero Value Config

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A service boots from a `Config` struct. Ops only ever want to pin the port in the
defaults — everything else should fall back to Go's zero values so unset fields
are unambiguous. A `Tags` slice that comes back as a non-nil empty slice, or a
`Host` pre-filled with a stray string, breaks the "unset means zero" contract the
rest of the system relies on.

## Task

Implement `DefaultConfig` in [zerovalues.go](zerovalues.go) so that it returns a
`Config` where:

1. `Port` is `8080`.
2. `Host` is `""`, `Debug` is `false`, `Tags` is `nil` — each left at its type's
   zero value, not explicitly re-set to look like it.
3. Each call returns a fresh value: mutating one result never affects the next.

Do **not** change the function signature or the tests.

> **Note — avoid the magic number.** Writing `Config{Port: 8080}` passes the
> tests, but hard-coding `8080` in the body is a *magic number*: an unexplained
> literal. Prefer naming it once — `const DefaultPort = 8080` — and returning
> `Config{Port: DefaultPort}`. The tests only grade the returned values, so this
> is style guidance, not a hard gate; get into the habit anyway.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DefaultConfig().Port
Output: 8080
```

**Example 2:**

```
Input:  DefaultConfig().Host
Output: ""
```

_Explanation:_ Unset string field keeps its zero value.

**Example 3:**

```
Input:  DefaultConfig().Tags
Output: nil
```

_Explanation:_ A nil slice, not an empty non-nil one.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Zero values** | A declared-but-unassigned variable gets its type's zero: `0`, `""`, `false`, `nil` slice/map/pointer. |
| 2 | **`var` vs composite literal** | `var c Config` gives an all-zero struct; a literal like `Config{Port: 8080}` sets named fields and leaves the rest zero. |
| 3 | **nil vs empty slice** | `[]string(nil)` and `[]string{}` both have length 0 but are not equal — the contract here wants nil. |

## Hint

You do not need to touch `Host`, `Debug`, or `Tags` at all — set only the field
that differs from zero. Reaching for `[]string{}` to "initialize" `Tags` is
exactly the mistake: that makes it a non-nil empty slice.

## Validate

```bash
make verify
```
