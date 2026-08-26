# JSON Unmarshal

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Continuing from `jsonmarshal`, the API needs to receive money amounts as
strings like `"$10.50"` and parse them back into `cents` internally.

## Task

Implement `UnmarshalJSON` on `*Money` in [jsonunmarshal.go](jsonunmarshal.go):

1. The `data` is a JSON string (it includes the `"` quotes).
2. Strip the quotes and the `$`.
3. Parse the dollars and cents into `m.Cents`.
4. Return an error if the format is invalid (e.g. missing `$`).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  json.Unmarshal(`"$10.50"`, &m)
Output: m.Cents == 1050
```

**Example 2:**

```
Input:  json.Unmarshal(`"10.50"`, &m)
Output: error (missing $)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **json.Unmarshaler interface** | `UnmarshalJSON([]byte) error` customizes JSON parsing. |
| 2 | **Pointer receiver** | Must mutate `m` so the caller sees the change. |
| 3 | **JSON strings** | `data` has quotes. Easiest is to `json.Unmarshal` into a Go string first. |

## Hint

First `json.Unmarshal(data, &s)` to get a Go string. Then check `strings.HasPrefix(s, "$")`.
Parse `s[1:]` (e.g. `fmt.Sscanf` or manually).

## Validate

```bash
make verify
```
