# Delete Map Entries

**Level:** middle  
**Topic:** 03-generics

## Context

A session store drops expired entries while a request handler is still reading the live map.

## Task

Implement the stub(s) in [mapsdeletefunc.go](mapsdeletefunc.go):

1. Implement `Prune` using `maps.Clone` and `maps.DeleteFunc`.
2. Leave the input untouched; return an empty (non-nil) map for empty or nil input.
3. The predicate sees both the key and the value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Prune({a:1,b:2}, dropEven)
Output: {a:1}
```

**Example 2:**

```
Input:  input after the call
Output: unchanged
```

**Example 3:**

```
Input:  Prune(nil, drop)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`maps.DeleteFunc`** | Mutates the map in place and returns nothing. |
| 2 | **Key and value predicate** | Deleting can depend on either half of the entry. |
| 3 | **Deleting while ranging** | Go allows it for maps — which is why this helper is safe. |

## Hint

`maps.DeleteFunc` returns nothing: it edits the map you give it.

## Validate

```bash
make verify
```
