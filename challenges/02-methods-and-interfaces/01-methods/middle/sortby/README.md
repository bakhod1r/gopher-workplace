# Sort By Field

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A dashboard sorts users by different criteria. The sort key is passed as a
method expression — the caller chooses `Person.ByAge`, `Person.ByName`, etc.

## Task

Implement `SortByField` in [sortby.go](sortby.go):

1. Sort `people` in place using `sort.Slice`.
2. Use `less` as the comparison function.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortByField([Charlie:30, Alice:20, Bob:25], Person.ByAge)
Output: [Alice:20, Bob:25, Charlie:30]
```

**Example 2:**

```
Input:  SortByField([A:1, B:2], Person.ByAge)
Output: [A:1, B:2]
```

**Example 3:**

```
Input:  SortByField([], Person.ByAge)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method expressions** | `Person.ByAge` is `func(Person, Person) bool`. |
| 2 | **sort.Slice** | `sort.Slice(s, func(i, j int) bool)` — bridge from index-based to value-based. |

## Hint

`sort.Slice(people, func(i, j int) bool { return less(people[i], people[j]) })`.

## Validate

```bash
make verify
```
