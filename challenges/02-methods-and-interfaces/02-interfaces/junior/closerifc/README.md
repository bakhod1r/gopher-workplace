# Closer Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A pipeline shuts resources down at the end of a run and reports the first failure.

## Task

Implement the stub(s) in [closerifc.go](closerifc.go):

1. Implement `Close` on `*File` — mark it closed and return an error if it was already closed.
2. Implement `CloseAll`, which closes every closer in order and returns the first error, or nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  f := &File{}; f.Close()
Output: nil
```

**Example 2:**

```
Input:  f := &File{}; f.Close(); f.Close()
Output: error "already closed"
```

**Example 3:**

```
Input:  CloseAll([]Closer{&File{}, &File{}})
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Small interfaces** | `Closer` is one method — the Go standard-library style. |
| 2 | **Pointer receiver mutation** | `Close` must change state, so the receiver is `*File`. |
| 3 | **Errors as values** | Reused: `errors.New`, and `nil` meaning success. |

## Hint

Loop, call `Close`, and return as soon as one call returns non-nil.

## Validate

```bash
make verify
```
