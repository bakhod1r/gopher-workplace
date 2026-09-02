# Decorator

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A pricing engine layers discounts and taxes on top of a base price calculator.

## Task

Implement the stub(s) in [decorator.go](decorator.go):

1. Implement `Price` on `Base`.
2. Implement `Price` on `Discount`, which takes `Percent` off the wrapped price (integer arithmetic, truncating).
3. Implement `Price` on `Tax`, which adds `Percent` to the wrapped price (truncating).
4. Implement `Wrap`, which layers the decorators so the first listed is applied to the base first.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Base{Amount: 100}.Price()
Output: 100
```

**Example 2:**

```
Input:  Discount{Inner: Base{100}, Percent: 10}.Price()
Output: 90
```

**Example 3:**

```
Input:  Tax{Inner: Discount{Base{100}, 10}, Percent: 20}.Price()
Output: 108
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Decorator pattern** | Each layer implements the same interface and holds the next one. |
| 2 | **Integer percentage maths** | Reused from data types: multiply before dividing. |
| 3 | **Composition over inheritance** | Behaviour is added by nesting, not by subclassing. |

## Hint

Discount: `p - p*Percent/100`. Tax: `p + p*Percent/100`.

## Validate

```bash
make verify
```
