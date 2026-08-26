# Functional Options Pattern

## Intuition

Instead of a constructor with many parameters (`NewServer(host, port, timeout,
maxConn, tls, ...)`), Go uses the functional options pattern: each option is a
small function that modifies one setting. This keeps the API clean and
backwards-compatible.

## Approach

1. Each `With*` returns `func(s *Server) { s.Field = val }`.
2. `NewServer` creates defaults, iterates options, applies each.

## Solution

```go
func WithHost(host string) Option {
	return func(s *Server) { s.Host = host }
}

func WithPort(port int) Option {
	return func(s *Server) { s.Port = port }
}

func WithTimeout(timeout int) Option {
	return func(s *Server) { s.Timeout = timeout }
}

func NewServer(opts ...Option) Server {
	s := Server{Host: "localhost", Port: 8080, Timeout: 30}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}
```

## Walkthrough

`NewServer(WithPort(9090))`:
1. `s = {localhost, 8080, 30}`.
2. `WithPort(9090)` → `func(s *Server) { s.Port = 9090 }`.
3. Apply: `s.Port = 9090`.
4. Return `{localhost, 9090, 30}`.

## Pitfalls

- Passing `s` (not `&s`) to options — the function takes `*Server`, so you
  need a pointer.
- Forgetting to set defaults — options only override, they don't create.
