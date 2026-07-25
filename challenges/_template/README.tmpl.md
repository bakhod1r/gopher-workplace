# {{TITLE}}

**Level:** {{LEVEL}}
**Topic:** {{TOPIC}} → {{SUBTOPIC}}
**Estimated time:** 15 min

## Context

<Short workplace story that motivates the bug. 2–4 sentences.>

## Task

Implement `{{FUNC}}` in [{{NAME}}.go]({{NAME}}.go) so that it:

1. <requirement>
2. <requirement>
3. <edge case: nil / empty / boundary>

Do **not** change the function signature or the tests.

## Examples

```go
{{FUNC}}(...) // => ...
{{FUNC}}(...) // => ...
{{FUNC}}(...) // => ...
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **<concept>** | <why it matters here> |
| 2 | **<concept>** | <why it matters here> |
| 3 | **<concept>** | <why it matters here> |

## Hint

<One paragraph pointing at the class of mistake, not the literal fix.>

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
