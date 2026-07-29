# CSV Field Quoting

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

An export endpoint writes CSV. Per RFC 4180 a field needs quoting only if it
contains a comma, quote, or newline — and inner quotes are doubled.

## Task

Implement `Quote(s)` following that rule.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Quote("plain")
Output: "plain"
```

_Explanation:_ no special chars -> unchanged

**Example 2:**

```
Input:  Quote("a,b")
Output: "\"a,b\""
```

_Explanation:_ comma forces quoting

**Example 3:**

```
Input:  Quote("say \"hi\"")
Output: "\"say \"\"hi\"\"\""
```

_Explanation:_ inner quotes doubled and field wrapped

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conditional quoting** | Only quote when a special char is present. |
| 2 | **Escaping** | Double each inner `"`. |
| 3 | **String scanning** | Detect specials with strings.ContainsAny or a loop. |

## Hint

If `strings.ContainsAny(s, ",\"\n")`, wrap in quotes and replace `"` with `""`;
else return `s`.

## Validate

```bash
make verify
```
