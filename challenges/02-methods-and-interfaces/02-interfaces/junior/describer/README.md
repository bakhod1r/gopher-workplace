# Describer

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An admin page prints a one-line description of every entity it lists.

## Task

Implement the stub(s) in [describer.go](describer.go):

1. Implement `Describe` on `User` — `"user <Name>"`.
2. Implement `Describe` on `Server` — `"server <Host>:<Port>"`.
3. Implement `DescribeAll`, which returns one description per element, in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  User{Name: "ann"}.Describe()
Output: "user ann"
```

**Example 2:**

```
Input:  Server{Host: "db", Port: 5432}.Describe()
Output: "server db:5432"
```

**Example 3:**

```
Input:  DescribeAll([]Describer{User{Name: "a"}})
Output: ["user a"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface satisfaction** | Two unrelated structs share one behaviour. |
| 2 | **fmt.Sprintf** | Reused from standard library basics: `%s` and `%d` verbs. |
| 3 | **Slice building** | Reused: `make` with a known length, or `append`. |

## Hint

`fmt.Sprintf("server %s:%d", s.Host, s.Port)`.

## Validate

```bash
make verify
```
