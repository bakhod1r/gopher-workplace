# Stringer

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Debug output prints domain values with `%v`, and the types decide how they look.

## Task

Implement the stub(s) in [stringer.go](stringer.go):

1. Implement `String` on `Color` — `"red"`, `"green"`, `"blue"`, or `"unknown"`.
2. Implement `String` on `Temp` — `"<n>C"`.
3. Implement `Print`, which returns the value's string form.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Red.String()
Output: "red"
```

**Example 2:**

```
Input:  Color(9).String()
Output: "unknown"
```

**Example 3:**

```
Input:  Print(Temp(21))
Output: "21C"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **fmt.Stringer** | The standard one-method interface `fmt` looks for when printing. |
| 2 | **Constant enums** | Reused from language basics: `iota`-style constants. |
| 3 | **switch on a value** | Reused: mapping an enum to text. |

## Hint

Implementing `String() string` makes `%v` and `%s` print your text automatically.

## Validate

```bash
make verify
```
