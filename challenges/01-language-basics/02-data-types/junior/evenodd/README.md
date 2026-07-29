# Integer Parity

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Parity is the classic use of the remainder operator `%` — but negative numbers
need care, because `-7 % 2` is `-1` in Go, not `1`.

## Task

Implement `Parity(n)` returning `"even"` or `"odd"`, correct for negatives.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Parity(0)
Output: "even"
```

_Explanation:_ 0%2==0.

**Example 2:**

```
Input:  Parity(3)
Output: "odd"
```

_Explanation:_ 3%2==1.

**Example 3:**

```
Input:  Parity(-4)
Output: "even"
```

_Explanation:_ -4%2==0 holds for negatives.

**Example 4:**

```
Input:  Parity(-7)
Output: "odd"
```

_Explanation:_ -7%2==-1 which is not 0, so odd. Testing ==1 would wrongly fail here.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Remainder `%`** | `n % 2` is 0 for even numbers. |
| 2 | **Sign of `%`** | In Go the result takes the sign of the dividend. |
| 3 | **Even test** | `n%2 == 0` works for negatives; `== 1` does not. |

## Hint

Test `n%2 == 0` for even — it holds for negatives. `n%2 == 1` fails for `-7`.

## Validate

```bash
make verify
```
