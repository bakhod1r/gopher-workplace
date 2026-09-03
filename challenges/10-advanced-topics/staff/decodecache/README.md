# Resolve The Layout Once Per Type, Safely

**Level:** staff
**Topic:** 10-advanced-topics / 03-reflection

## Context

A decoder walks the struct's field table on every call. It shows up in the profile of a service that decodes the same three config types a hundred thousand times a second.

## Task

Implement [decodecache.go](decodecache.go):

1. Fill `dst`'s exported string fields from `src`, matching by `env` tag.
2. Use the shared `layoutOf` cache so a type's field table is walked once.
3. Leave untagged, unexported, non-string fields and missing keys alone.
4. Return `ErrTarget` unless `dst` is a non-nil pointer to a struct.
5. Correct under concurrent use — many goroutines decoding at once.

Replace the stub body in [decodecache.go](decodecache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Decode(map[string]string{"HOST":"h"}, &cfg)
Output: nil, cfg.Host is "h"
```

**Example 2:**

```
Input:  Decode(src, cfg{})
Output: ErrTarget
```

**Example 3:**

```
Input:  16 goroutines x 200 decodes
Output: every result correct
```

_Explanation:_ The cache is shared; the targets are not.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type metadata is immutable** | A resolved layout is valid forever, which is what makes caching sound. |
| 2 | **sync.Map for read-heavy caches** | Loads after the first write take no lock. |
| 3 | **LoadOrStore races benignly** | Two goroutines may both compute the layout; only one is published, and both are equal. |
| 4 | **Shared cache, private targets** | The cache is read-only after publication; each caller writes its own struct. |

## Hint

The cache is already written for you. The body is a validate, a lookup and a loop over the layout.

## Validate

```bash
make verify
```
