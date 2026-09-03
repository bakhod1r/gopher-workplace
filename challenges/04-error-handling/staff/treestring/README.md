# Render The Tree

**Level:** staff
**Topic:** 04-error-handling

## Context

A debugging endpoint prints an error tree with indentation so an engineer can see which failure came from which branch.

## Task

Implement `Tree` in [treestring.go](treestring.go):

1. Render one line per error, indented by one tab per level.
2. Render children after their parent, joined branches in order.
3. Return `""` for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Tree(ErrA)
Output: "a"
```

**Example 2:**

```
Input:  Tree(fmt.Errorf("x: %w", ErrA))
Output: "x: a\n\ta"
```

**Example 3:**

```
Input:  Tree(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive rendering** | Depth becomes indentation. |
| 2 | **Pre-order output** | Parent line precedes its children. |
| 3 | **strings.Repeat** | Building the indent prefix. |

## Hint

Pass the depth down the recursion; the root is depth 0 and gets no indent.

## Validate

```bash
make verify
```
