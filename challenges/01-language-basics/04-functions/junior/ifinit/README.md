# If with Init Clause

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

An `if` init clause (`if r := n % 3; ...`) computes a value once and scopes it to the if/else chain.

## Task

Implement `Bucket` in [ifinit.go](ifinit.go) classifying `n` by `n % 3`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bucket(9)
Output: "zero"
```

**Example 2:**

```
Input:  Bucket(10)
Output: "one"
```

**Example 3:**

```
Input:  Bucket(11)
Output: "two"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if init clause** | `if r := expr; cond` scopes `r`. |
| 2 | **modulo** | `n % 3` gives 0,1,2. |
| 3 | **scoped variable** | `r` is visible only in the if/else. |

## Hint

`if r := n % 3; r == 0 { ... } else if r == 1 { ... }`.

## Validate

```bash
make verify
```
