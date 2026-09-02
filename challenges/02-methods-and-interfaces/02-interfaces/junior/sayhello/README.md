# Say Hello

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A multilingual greeter picks the language from the visitor's locale object.

## Task

Implement the stub(s) in [sayhello.go](sayhello.go):

1. Implement `Hello` on `English` and `Uzbek`.
2. Implement `Greet`, which returns the greeting followed by `", <name>"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  English{}.Hello()
Output: "Hello"
```

**Example 2:**

```
Input:  Uzbek{}.Hello()
Output: "Salom"
```

**Example 3:**

```
Input:  Greet(Uzbek{}, "Ali")
Output: "Salom, Ali"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface dispatch** | One call site, many greetings. |
| 2 | **Empty struct implementers** | Behaviour without state. |
| 3 | **String concatenation** | Reused: building the final line. |

## Hint

`return g.Hello() + ", " + name`.

## Validate

```bash
make verify
```
