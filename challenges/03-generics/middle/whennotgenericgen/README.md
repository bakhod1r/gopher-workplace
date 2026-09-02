# When Not To Use Generics

**Level:** middle  
**Topic:** 03-generics

## Context

A code review asks whether a helper should take an interface or a type parameter. Both appear here so the difference is concrete.

## Task

Implement the stub(s) in [whennotgenericgen.go](whennotgenericgen.go):

1. Implement `WriteAll` and `WriteEach`, each returning the number of writes.
2. Note the rule: one value needing behaviour takes an interface; a typed collection takes a type parameter.
3. Neither function should know which concrete writer it received.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WriteAll(w, []string{"a","b"})
Output: 2
```

**Example 2:**

```
Input:  WriteEach(w, []int{1,2}, itoa)
Output: 2
```

**Example 3:**

```
Input:  WriteAll(w, nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interfaces for behaviour** | One value, one method set — a type parameter would add nothing. |
| 2 | **Type parameters for collections** | They avoid boxing every element and keep the caller's slice typed. |
| 3 | **Cost of over-generalising** | A type parameter with a single instantiation is complexity for nothing. |

## Hint

`w` is one value used for its behaviour — that is an interface, not a type parameter.

## Validate

```bash
make verify
```
