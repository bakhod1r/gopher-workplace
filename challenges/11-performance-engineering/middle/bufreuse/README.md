# One Buffer, Many Records

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

An encoder that returns a fresh `[]byte` per record allocates once per record — millions of short-lived objects for the collector to sweep. Owning one buffer and handing out a view of it removes every one of those allocations, at the price of a contract the caller must know about: the bytes are yours until the next call.

## Task

Implement both methods in [bufreuse.go](bufreuse.go):

1. `Encode` renders `"name=value;"` for each pair into the encoder's own buffer and returns it, pairing up to the shorter slice.
2. A warm `Encode` must allocate nothing, and each call replaces the previous record rather than appending to it.
3. `Clone` returns a copy that survives the next `Encode`.

## Examples

**Example 1:**

```
Input:  Encode([a b], [1 2])
Output: "a=1;b=2;"
```

**Example 2:**

```
Input:  Encode([a b], [1])
Output: "a=1;"
```

**Example 3:**

```
Input:  Encode(a); saved := Clone(); Encode(b)
Output: saved is still "a=1;"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reuse means aliasing** | The returned slice shares the encoder's array, so its lifetime is bounded. |
| 2 | **Reset with `buf[:0]`** | Length zero keeps the array; forgetting it concatenates every record ever encoded. |
| 3 | **Offer an escape hatch** | `Clone` lets callers who need to keep the data pay for it explicitly. |

## Topics used again

`append`, slice reuse, `min`, `slices.Clone`, methods on pointer receivers.

## Hint

Reset at the top of `Encode`, then append; the buffer grows once and then stops.

## Validate

```bash
make verify
```
