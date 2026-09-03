# The Append API Shape

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`strconv.AppendInt`, `time.Time.AppendFormat`, `fmt.Appendf` all share one signature: take the caller's buffer, append, return it. The caller owns the memory and can reuse it forever, so a hot loop that formats millions of records allocates nothing at all.

## Task

Implement `AppendRecord` in [byteappend.go](byteappend.go):

1. Append `key`, `=`, `value`, `;` to `dst` and return the extended slice.
2. Never allocate when `dst` already has spare capacity.
3. A nil `dst` still works, producing just the record.

## Examples

**Example 1:**

```
Input:  AppendRecord(nil, "a", "1")
Output: "a=1;"
```

**Example 2:**

```
Input:  AppendRecord([]byte("x:"), "k", "v")
Output: "x:k=v;"
```

**Example 3:**

```
Input:  AppendRecord(nil, "", "")
Output: "=;"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Caller-owned buffers** | The allocation happens once, outside the loop, under the caller's control. |
| 2 | **`append(b, s...)`** | Appending a string to a `[]byte` needs no conversion and no copy. |
| 3 | **Return the result** | `append` may move the data, so the returned slice is the only valid one. |

## Topics used again

`append`, slices, string-to-byte-slice appends.

## Hint

Three appends and one return; no `make` anywhere in the function.

## Validate

```bash
make verify
```
