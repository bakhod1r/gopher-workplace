# Let The Caller Own The Memory

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

The standard library's `Append*` family exists because the alternative — returning a freshly allocated slice — forces an allocation the caller often does not need. Hand the function a scratch buffer, get the result appended, reuse the buffer forever. The companion sizing function is what lets the caller allocate exactly once.

## Task

Implement both functions in [scratchbuf.go](scratchbuf.go):

1. `AppendJoin` appends the parts separated by `sep` to `scratch` and returns the extended slice, preserving whatever was already there.
2. It must not allocate when `scratch` has room.
3. `Sized` returns exactly how many bytes `AppendJoin` will append.

## Examples

**Example 1:**

```
Input:  AppendJoin(nil, [a b], "/")
Output: "a/b"
```

**Example 2:**

```
Input:  AppendJoin([]byte("prefix:"), [a b], "/")
Output: "prefix:a/b"
```

**Example 3:**

```
Input:  Sized([ab cd ef], "--")
Output: 10
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Append APIs move the allocation** | The caller decides where the memory comes from and how long it lives. |
| 2 | **A sizing function completes the pattern** | Without it the caller has to guess, and guessing means regrowing. |
| 3 | **`n-1` separators** | The classic off-by-one in every join implementation. |

## Topics used again

`append` with strings, capacity reuse, integer arithmetic.

## Hint

`Sized` is the same arithmetic the caller would need to `Grow` a `strings.Builder`.

## Validate

```bash
make verify
```
