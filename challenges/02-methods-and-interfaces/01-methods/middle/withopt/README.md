# Functional Options

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A server factory uses the **functional options** pattern — a Go idiom for
clean, extensible configuration without giant constructors.

## Task

Implement `WithHost`, `WithPort`, `WithTimeout`, and `NewServer` in
[withopt.go](withopt.go):

1. Each `With*` returns an `Option` (a `func(*Server)`) that sets one field.
2. `NewServer` creates a server with defaults, applies all options, returns it.
3. Defaults: `Host="localhost"`, `Port=8080`, `Timeout=30`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewServer()
Output: {Host:"localhost", Port:8080, Timeout:30}
```

**Example 2:**

```
Input:  NewServer(WithPort(9090))
Output: {Host:"localhost", Port:9090, Timeout:30}
```

**Example 3:**

```
Input:  NewServer(WithHost("0.0.0.0"), WithPort(3000))
Output: {Host:"0.0.0.0", Port:3000, Timeout:30}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Functional options** | `type Option func(*Server)` — function as configuration unit. |
| 2 | **Closures** | Each `With*` returns a closure capturing its argument. |
| 3 | **Variadic functions** | `opts ...Option` accepts zero or more options. |
| 4 | **Pointer receiver** | Options take `*Server` to mutate the defaults. |

## Hint

Each `With*` returns `func(s *Server) { s.Field = value }`.
`NewServer` creates defaults, loops over `opts`, calls each on `&s`.

## Validate

```bash
make verify
```
