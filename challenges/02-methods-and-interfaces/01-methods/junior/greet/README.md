# Greet

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A chat application needs to display a personalised greeting when a user
connects. The greeting is built from the user's name.

## Task

Implement the `Greet` method on `Person` in [greet.go](greet.go) so that:

1. It returns `"Hello, <Name>!"` where `<Name>` is the person's `Name` field.
2. If `Name` is empty, it still returns `"Hello, !"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Person{Name: "Alice"}.Greet()
Output: "Hello, Alice!"
```

**Example 2:**

```
Input:  Person{Name: "Bob"}.Greet()
Output: "Hello, Bob!"
```

**Example 3:**

```
Input:  Person{Name: ""}.Greet()
Output: "Hello, !"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods vs Functions** | A method has a receiver — `func (p Person) Greet()` ties `Greet` to `Person`. |
| 2 | **Value receiver** | `Person` is small; a value receiver copies it — fine for read-only access. |
| 3 | **String concatenation** | Build the greeting with `+` or `fmt.Sprintf`. |

## Hint

A method is just a function whose first parameter is the receiver.
Use `"Hello, " + p.Name + "!"` or `fmt.Sprintf("Hello, %s!", p.Name)`.

## Validate

```bash
make verify
```
