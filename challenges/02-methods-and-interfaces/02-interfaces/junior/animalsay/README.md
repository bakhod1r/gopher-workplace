# Animal Say

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A children's app plays the sound of whichever animal is tapped.

## Task

Implement the stub(s) in [animalsay.go](animalsay.go):

1. Implement `Sound` on `Dog` (`"Woof!"`) and `Cat` (`"Meow!"`).
2. Implement `MakeNoise`, which returns the animal's sound.
3. Implement `Chorus`, which joins every animal's sound with a space.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MakeNoise(Dog{})
Output: "Woof!"
```

**Example 2:**

```
Input:  MakeNoise(Cat{})
Output: "Meow!"
```

**Example 3:**

```
Input:  Chorus([]Animal{Dog{}, Cat{}})
Output: "Woof! Meow!"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface satisfaction** | No declaration is needed — having the method is enough. |
| 2 | **Polymorphic call** | `MakeNoise` works for any animal, present or future. |
| 3 | **Joining strings** | Reused: separator only between elements. |

## Hint

In `Chorus`, add the space before every sound except the first.

## Validate

```bash
make verify
```
