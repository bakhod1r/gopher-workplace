# Fill A Struct From A String Map

**Level:** middle
**Topic:** 10-advanced-topics / 03-reflection

## Context

Every service in the fleet parses its own environment by hand. The parsing is identical, the mistakes are not, and each one is found in production.

## Task

Implement [tagdecode.go](tagdecode.go):

1. Fill `dst`'s fields from `src`, matching each field's `env` tag against the map key.
2. Support string, int and bool fields; report an error for any other tagged kind.
3. Skip unexported fields, untagged fields, `env:"-"`, and keys absent from `src`.
4. Return `ErrTarget` unless `dst` is a non-nil pointer to a struct.

Replace the stub body in [tagdecode.go](tagdecode.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Decode(map[string]string{"PORT":"8080"}, &cfg)
Output: nil, cfg.Port is 8080
```

**Example 2:**

```
Input:  Decode(map[string]string{"PORT":"eighty"}, &cfg)
Output: a parse error
```

**Example 3:**

```
Input:  Decode(src, config{})
Output: ErrTarget
```

_Explanation:_ A value cannot be written through.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Tag-driven mapping** | The tag is the contract between the external name and the field. |
| 2 | **Kind switching** | Each supported kind needs its own parse and its own `Set` method. |
| 3 | **Settability** | The pointer plus `Elem` is what makes the fields writable. |
| 4 | **Error wrapping** | `%w` keeps the underlying `strconv` error inspectable. |

## Hint

One loop over the fields. Per field: tag, lookup, kind switch, set.

## Validate

```bash
make verify
```
