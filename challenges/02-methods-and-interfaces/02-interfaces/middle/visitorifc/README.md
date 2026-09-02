# Visitor

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A document tree is walked by different visitors: one counts words, another collects headings.

## Task

Implement the stub(s) in [visitorifc.go](visitorifc.go):

1. Implement `Accept` on `Text` and `Section` — `Section` visits itself, then every child in order.
2. Implement `Visit` on `*WordCounter` (count words in text nodes) and `*HeadingCollector` (collect section titles).
3. Implement `Walk`, which runs a visitor over a node.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  word count of Section{Children: [Text{"a b"}]}
Output: 2
```

**Example 2:**

```
Input:  headings of a two-level section tree
Output: outer then inner title
```

**Example 3:**

```
Input:  visiting a bare Text node
Output: only that node is visited
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Visitor pattern** | Node types and operations vary independently. |
| 2 | **Double dispatch** | `Accept` picks the node; `Visit` picks the operation. |
| 3 | **Recursive traversal** | Reused: pre-order walk over a tree. |

## Hint

`Section.Accept` calls `v.Visit(s)` first, then `child.Accept(v)` for each child.

## Validate

```bash
make verify
```
