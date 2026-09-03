# Look Up Without Copying, Store With A Copy

**Level:** staff
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A counter over a hot byte stream copies every key twice: once to look it up and once again when it is already present. Ninety-nine percent of the keys are already present.

## Task

Implement [zerocopykeys.go](zerocopykeys.go):

1. Increment `m`'s counter for each key.
2. An existing key must be counted without allocating.
3. A key stored for the first time must own its bytes — the caller reuses the buffers.
4. Empty and nil keys are skipped.

Replace the stub body in [zerocopykeys.go](zerocopykeys.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Count(m, [][]byte{[]byte("a"), []byte("a")})
Output: m[a] == 2
```

**Example 2:**

```
Input:  200 counts of an existing key
Output: 0 allocations
```

**Example 3:**

```
Input:  26 batches through one reused buffer
Output: 26 distinct keys
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Borrowing for a lookup** | The borrowed string dies inside the call, so nothing can observe the aliasing. |
| 2 | **Owning for a store** | A stored key is read again later; it must not be able to change. |
| 3 | **Map keys cache their hash** | A key whose bytes change strands its entry permanently. |
| 4 | **The asymmetry is the whole design** | Read paths may alias; write paths may not. |

## Hint

Two paths through the loop. Only one of them may allocate.

## Validate

```bash
make verify
```
