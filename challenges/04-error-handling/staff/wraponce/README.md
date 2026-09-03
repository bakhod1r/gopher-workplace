# Idempotent Annotation

**Level:** staff
**Topic:** 04-error-handling

## Context

Two layers both prefix an error with the same operation name, producing `"save: save: …"`. The wrapper must add the prefix only when it is not already there.

## Task

Implement `Once` in [wraponce.go](wraponce.go):

1. Return nil for a nil error.
2. Wrap `err` as `"<op>: <err>"` when its message does not already start with that prefix.
3. Return `err` unchanged when the prefix is already present.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Once("save", ErrA)
Output: "save: a"
```

**Example 2:**

```
Input:  Once("save", Once("save", ErrA))
Output: "save: a"
```

**Example 3:**

```
Input:  Once("save", nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Idempotence** | Applying twice equals applying once. |
| 2 | **strings.HasPrefix** | Cheap check before allocating. |
| 3 | **Identity on skip** | The same value is returned, not a copy. |

## Hint

The check is on the message; the skip must return the identical error value.

## Validate

```bash
make verify
```
