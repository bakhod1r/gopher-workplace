# Speaker

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A demo app makes every character on screen introduce itself.

## Task

Implement the stub(s) in [speaker.go](speaker.go):

1. Implement `Speak` on `Person` — `"Hi, I'm <Name>"`.
2. Implement `Speak` on `Robot` — always `"I am robot"`.
3. Implement `Introduce`, which returns what the speaker says.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Person{Name: "Go"}.Speak()
Output: "Hi, I'm Go"
```

**Example 2:**

```
Input:  Robot{ID: 1}.Speak()
Output: "I am robot"
```

**Example 3:**

```
Input:  Introduce(Person{Name: "Ann"})
Output: "Hi, I'm Ann"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface variable** | `var s Speaker` can hold either concrete type in turn. |
| 2 | **Ignoring receiver fields** | `Robot` has an `ID` its method never reads — that is allowed. |
| 3 | **String literals with quotes** | Reused: an apostrophe inside a double-quoted string. |

## Hint

`"Hi, I'm " + p.Name` — the apostrophe needs no escaping in a double-quoted string.

## Validate

```bash
make verify
```
