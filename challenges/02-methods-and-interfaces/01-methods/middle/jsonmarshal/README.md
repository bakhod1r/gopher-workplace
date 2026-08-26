# JSON Marshal

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An API returns money amounts. Internally they're stored as cents (int), but the
JSON output should be a human-readable dollar string: `"$10.50"`.

## Task

Implement `MarshalJSON` on `Money` in [jsonmarshal.go](jsonmarshal.go):

1. Format cents as `"$X.YY"` (always two decimal places).
2. Return as JSON string bytes (with quotes).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Money{1050}.MarshalJSON()
Output: []byte(`"$10.50"`), nil
```

**Example 2:**

```
Input:  Money{99}.MarshalJSON()
Output: []byte(`"$0.99"`), nil
```

**Example 3:**

```
Input:  Money{0}.MarshalJSON()
Output: []byte(`"$0.00"`), nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **json.Marshaler interface** | `MarshalJSON() ([]byte, error)` customizes JSON output. |
| 2 | **Value receiver** | `Money` is small — copy is fine. |
| 3 | **fmt.Sprintf** | `"$%d.%02d"` formats dollars and cents. |

## Hint

`dollars := m.Cents / 100`, `cents := m.Cents % 100`.
Return `[]byte(fmt.Sprintf(`"$%d.%02d"`, dollars, cents))`.

## Validate

```bash
make verify
```
