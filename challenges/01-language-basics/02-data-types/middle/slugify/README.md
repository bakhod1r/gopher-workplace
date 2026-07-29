# URL Slug

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A CMS turns article titles into URL slugs: lowercase, alphanumerics kept,
everything else collapsed to single dashes, edges trimmed.

## Task

Implement `Slug(s)` per the rules above.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Slug("Hello, World!")
Output: "hello-world"
```

_Explanation:_ lowercased, punctuation run -> single dash, trailing '!' trimmed

**Example 2:**

```
Input:  Slug("  Go 1.26  ")
Output: "go-1-26"
```

_Explanation:_ leading/trailing spaces trimmed, '.' becomes dash

**Example 3:**

```
Input:  Slug("a---b")
Output: "a-b"
```

_Explanation:_ run of non-alnum collapses to one dash

**Example 4:**

```
Input:  Slug("!!!")
Output: ""
```

_Explanation:_ nothing alphanumeric -> empty

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune classification** | Keep `0-9a-z`, lowercase letters. |
| 2 | **Run collapsing** | Emit one `-` per run of separators. |
| 3 | **Trim edges** | No leading/trailing `-`. |

## Hint

Scan runes; append lowercased alnum; for a separator set a "pending dash" flag
that emits a single `-` before the next alnum.

## Validate

```bash
make verify
```
