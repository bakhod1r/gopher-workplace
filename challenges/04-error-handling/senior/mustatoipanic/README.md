# Strict Parse

**Level:** senior
**Topic:** 04-error-handling

## Context

A code generator parses constants it produced itself. A malformed value there means the generator is broken, so parsing fails loudly.

## Task

Implement `MustParse` in [mustatoipanic.go](mustatoipanic.go):

1. Return the parsed integer on success.
2. Panic with an error wrapping the parse failure and quoting the input.
3. Keep `strconv.ErrSyntax` matchable in the panicked error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MustParse("42")
Output: 42
```

**Example 2:**

```
Input:  MustParse("x")
Output: panics
```

**Example 3:**

```
Input:  errors.Is(recovered, strconv.ErrSyntax)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Panic payload as error** | Recoverable and inspectable. |
| 2 | **Wrapping before panicking** | Context survives the panic. |
| 3 | **Must semantics** | Only for inputs the program controls. |

## Hint

Build the wrapped error first, then panic with it — the test recovers and calls `errors.Is` on the payload.

## Validate

```bash
make verify
```
