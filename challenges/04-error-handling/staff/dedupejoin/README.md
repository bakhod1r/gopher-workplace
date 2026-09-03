# Join Without Repeats

**Level:** staff
**Topic:** 04-error-handling

## Context

Every replica reports the same downstream outage. The combined error should say it once.

## Task

Implement `Distinct` in [dedupejoin.go](dedupejoin.go):

1. Join the non-nil errors, keeping only the first with each message.
2. Preserve first-seen order.
3. Return nil when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Distinct(ErrA, ErrA, ErrB)
Output: matches ErrA and ErrB, message has two lines
```

**Example 2:**

```
Input:  Distinct(nil)
Output: nil
```

**Example 3:**

```
Input:  Distinct(ErrA)
Output: matches ErrA
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deduplication before joining** | The join reflects what is kept. |
| 2 | **First occurrence wins** | Identity of the representative matters. |
| 3 | **errors.Join semantics** | Nil entries are dropped for you. |

## Hint

Deduplicate first, then join once — joining first leaves the duplicates inside the tree.

## Validate

```bash
make verify
```
