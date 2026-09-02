# Prefix Tree

**Level:** middle  
**Topic:** 03-generics

## Context

An autocomplete box answers "does anything start with this?" on every keystroke, so a full scan of the dictionary is out.

## Task

Implement the stub(s) in [triegen.go](triegen.go):

1. Implement `NewTrie`, `Insert`, `Get`, and `HasPrefix`.
2. A key stored with a zero value must still be found by `Get`.
3. `HasPrefix("")` is true for any trie, including an empty one.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Insert("go", 1); Get("go")
Output: 1, true
```

**Example 2:**

```
Input:  Get("g")
Output: zero, false
```

**Example 3:**

```
Input:  Insert("go", 1); HasPrefix("g")
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive generic types** | Each node maps a rune to another `*Trie[V]` — the same instantiation. |
| 2 | **Presence flags** | A node on the path is not necessarily a stored key; `set` records which nodes are. |
| 3 | **Ranging a string** | `for _, r := range key` walks runes, not bytes. |

## Hint

A node existing on the path does not mean a value was stored there — that is what `set` is for.

## Validate

```bash
make verify
```
