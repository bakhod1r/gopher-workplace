# Custom As Method

**Level:** staff
**Topic:** 04-error-handling

## Context

A legacy error type must be convertible into the modern structured type so callers can migrate without a flag day.

## Task

Implement `Modern` in [customas.go](customas.go):

1. Give `*LegacyError` an `As(target any) bool` that fills a `**Modern` with an equivalent value.
2. Copy `Op` and map `Num` onto `Code`.
3. Return false for any other target type.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  errors.As(&LegacyError{"read", 5}, &m)
Output: true, m.Code == 5
```

**Example 2:**

```
Input:  errors.As(legacy, &other)
Output: false
```

**Example 3:**

```
Input:  errors.As(wrapped, &m)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The As method** | `errors.As` consults it before matching types. |
| 2 | **Adapting types** | One error can present itself as another. |
| 3 | **Target discipline** | The target is always a pointer to the wanted type. |

## Hint

The argument is `any` holding a `**Modern` — assert to that, then assign through it.

## Validate

```bash
make verify
```
