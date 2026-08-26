# Method Values as HTTP Handlers

## Intuition

`http.HandlerFunc` is `type HandlerFunc func(ResponseWriter, *Request)`. A
method `(a *App) HealthHandler(w, r)` has exactly that signature when bound as
a method value. No wrapper needed.

## Approach

1. Return `a.HealthHandler`.

## Solution

```go
func (a *App) Handler() http.HandlerFunc {
	return a.HealthHandler
}
```

## Walkthrough

- `a.HealthHandler` binds `a` (pointer receiver).
- Returns `func(http.ResponseWriter, *http.Request)`.
- `http.HandlerFunc` accepts this signature.

## Pitfalls

- Wrapping: `func(w, r) { a.HealthHandler(w, r) }` works but is unnecessary.
- Using value receiver would copy App — pointer ensures shared state.
