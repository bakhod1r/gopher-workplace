# Swimmer

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An aquarium screen lists what each creature is doing right now.

## Task

Implement the stub(s) in [swimmer.go](swimmer.go):

1. Implement `Swim` on `Fish` — `"<Name> swims"`.
2. Implement `Swim` on `Duck` — `"duck swims"`.
3. Implement `SwimAll`, which returns one line per swimmer, in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fish{Name: "nemo"}.Swim()
Output: "nemo swims"
```

**Example 2:**

```
Input:  Duck{}.Swim()
Output: "duck swims"
```

**Example 3:**

```
Input:  SwimAll([]Swimmer{Fish{Name: "a"}, Duck{}})
Output: ["a swims", "duck swims"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface satisfaction** | A struct with fields and an empty struct implement the same contract. |
| 2 | **Interface variable reassignment** | One variable, different dynamic types over time. |
| 3 | **Slice of results** | Reused: `append` into a preallocated slice. |

## Hint

`f.Name + " swims"`.

## Validate

```bash
make verify
```
