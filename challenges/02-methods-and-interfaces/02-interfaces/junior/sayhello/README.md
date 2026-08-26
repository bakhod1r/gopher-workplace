# Say Hello

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Different types satisfy the same interface by implementing its methods. The
caller doesn't need to know the concrete type.

## Task

Implement `Greet` on `English` and `Uzbek` in [sayhello.go](sayhello.go):

1. `English.Greet()` returns `"Hello!"`.
2. `Uzbek.Greet()` returns `"Salom!"`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  SayHello(English{})
Output: "Hello!"
```

**Example 2:**

```
Input:  SayHello(Uzbek{})
Output: "Salom!"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface definition** | `type Greeter interface { Greet() string }` declares a contract. |
| 2 | **Implicit satisfaction** | A type satisfies an interface by having all its methods. |

## Validate

```bash
make verify
```
