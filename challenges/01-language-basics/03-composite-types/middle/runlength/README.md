# Run-Length Encode

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A simple compression: replace runs of a repeated character with the character
and a count.

## Task

Implement `Encode(s)` producing `char + count` per run.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "aaab"
Output: "a3b1"
```

**Example 2:**

```
Input:  "abc"
Output: "a1b1c1"
```

**Example 3:**

```
Input:  ""
Output: ""
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Run detection** | Count while the byte repeats. |
| 2 | **Builder + Itoa** | Assemble char and number. |
| 3 | **Empty input** | Return "". |

## Hint

Walk bytes; count the current run; when it ends, write byte + `strconv.Itoa(n)`.

## Validate

```bash
make verify
```
