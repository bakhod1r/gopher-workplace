# Must That Never Panics

**Level:** staff  
**Topic:** 03-generics

## Context

A service wraps its start-up config load in `Must`. When the config file is missing the process starts anyway with an empty config and serves traffic against no upstreams at all.

## Task

Fix the single planted bug in [mustswallowbug.go](mustswallowbug.go):

1. Find and fix the single bug so a non-nil error aborts instead of returning a value.
2. The success path must still return `v` unchanged.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Must(7, nil)
Output: 7
```

**Example 2:**

```
Input:  Must(0, errors.New("boom"))
Output: panics
```

**Example 3:**

```
Input:  Must("x", nil)
Output: "x"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Silent zero values** | Returning the zero value on an error path hides the failure at the call site. |
| 2 | **`Must` is a contract** | The whole point of the helper is that the failure path is loud and terminal. |
| 3 | **Generic zero values** | `var zero T` compiles for every `T`, which is what makes this mistake so easy to hide. |

## Hint

What does the caller see when `err` is non-nil?

## Validate

```bash
make verify
```
