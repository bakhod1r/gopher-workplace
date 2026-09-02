# Three Stage CSV Import

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Customer CSV uploads arrive with padded cells and blank lines. The importer
runs three stages — read, clean, validate — each in its own goroutine, so a
slow cleaning step never stops the reader from getting ahead. The accepted
rows must come out in file order for the error report line numbers to mean
anything.

## Task

Implement `ImportCSV` in [csvimport.go](csvimport.go) so that:

1. A reader goroutine streams `rows` on a channel and closes it.
2. A cleaning goroutine forwards `clean(row)` for each row on its own channel.
3. A validation goroutine forwards only the rows where `valid(row)` holds; the caller collects them in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ImportCSV([]string{" a ", "  "}, TrimSpace, nonEmpty)
Output: []string{"a"}
```

**Example 2:**

```
Input:  ImportCSV([]string{" z", "", "a "}, TrimSpace, nonEmpty)
Output: []string{"z", "a"}
```

**Example 3:**

```
Input:  ImportCSV(nil, TrimSpace, nonEmpty)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Multi-stage pipeline** | Three goroutines, three channels, one direction of flow. |
| 2 | **Cascading close** | Each stage's close ends the next stage's range, all the way to the caller. |
| 3 | **Order preservation** | A single-channel chain never reorders, however many stages it has. |

## Hint

Every stage is the same four lines: make a channel, start a goroutine with
`defer close`, range over the previous channel, send onward.

## Validate

```bash
make verify
```
