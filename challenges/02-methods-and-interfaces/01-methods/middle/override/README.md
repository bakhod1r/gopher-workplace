# Override

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A pet store displays animal info. `Animal` has a `String()` method, but `Pet`
needs a richer version that includes the pet's nickname.

## Task

Implement `String` on `Pet` in [override.go](override.go):

1. Return `"Pet(<Nickname>, <Species>)"`.
2. This **shadows** the promoted `Animal.String()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Pet{Animal{"Cat"}, "Whiskers"}.String()
Output: "Pet(Whiskers, Cat)"
```

**Example 2:**

```
Input:  Pet{Animal{"Dog"}, "Rex"}.String()
Output: "Pet(Rex, Dog)"
```

**Example 3:**

```
Input:  Pet{Animal{""}, ""}.String()
Output: "Pet(, )"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method shadowing** | Defining `String` on `Pet` hides the promoted `Animal.String`. |
| 2 | **Accessing shadowed method** | `p.Animal.String()` still works if you need the original. |
| 3 | **fmt.Stringer** | `String()` is called by `fmt.Print` and friends. |

## Hint

`fmt.Sprintf("Pet(%s, %s)", p.Nickname, p.Species)` — `Species` is still
accessible via promotion.

## Validate

```bash
make verify
```
