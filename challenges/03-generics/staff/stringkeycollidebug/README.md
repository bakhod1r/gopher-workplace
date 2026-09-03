# Keys That Print The Same

**Level:** staff  
**Topic:** 03-generics

## Context

A deduplication pass over a heterogeneous event stream is collapsing records that
have nothing in common beyond how they render: the integer `1` and the string
`"1"` become one event, and two struct types with identical fields become one
row. The stream is also 200k events per batch, and the pass has become the
slowest stage in the pipeline.

## Task

Fix the single planted bug in [stringkeycollidebug.go](stringkeycollidebug.go):

1. Two elements are duplicates only when they are equal under `==`.
2. Elements that merely share a textual rendering must both survive.
3. Input order must be preserved, keeping the first of each duplicate group.
4. The pass must stay linear — no per-element formatting.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Distinct([]any{1, 1, 2})
Output: []any{1, 2}
```

**Example 2:**

```
Input:  Distinct([]any{1, "1"})
Output: []any{1, "1"}
```

**Example 3:**

```
Input:  Distinct([]string{"a", "b", "a"})
Output: []string{"a", "b"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`comparable` is the whole point** | The type parameter already guarantees `==` and map-key usability; nothing needs to be derived. |
| 2 | **Rendering is not identity** | `fmt.Sprint` erases the dynamic type, so values of different types collapse onto one key. |
| 3 | **Interface equality** | For an `any`, `==` compares the dynamic type *and* the value; `1` and `"1"` are never equal. |
| 4 | **Formatting is not free** | `fmt.Sprint` allocates and reflects per call, so a stringly-keyed set is far slower than the direct one. |

## Hint

`T` is already `comparable`. What does the code use as the map key instead?

## Validate

```bash
make verify
```
