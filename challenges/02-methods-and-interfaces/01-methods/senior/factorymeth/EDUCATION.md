# Factory Method

## Intuition

A constructor that always builds the same type is just a function. A factory
becomes interesting when the *type* of the result is a runtime decision: the
return type is an interface, and the concrete type behind it varies.

## Approach

1. Switch on the key.
2. Return the matching concrete value — it converts to `Store` implicitly.
3. Return `nil` for anything unrecognised.

## Solution

```go
func (f StoreFactory) Create(storeType string) Store {
	switch storeType {
	case "mem":
		return MemStore{}
	case "disk":
		return DiskStore{}
	default:
		return nil
	}
}
```

## Walkthrough

`MemStore` declares `Save` with a value receiver, so both `MemStore` and
`*MemStore` satisfy `Store`. Returning the value keeps the assertion
`.(MemStore)` in the test true.

The `default` arm returns an untyped `nil`, so the interface value is nil in
both its type and value words — which is why `f.Create("unknown") != nil` is
false.

## Pitfalls

- **Returning `&MemStore{}`.** Still a valid `Store`, but the test's
  `.(MemStore)` assertion fails: `*MemStore` is a different dynamic type.
- **The typed-nil trap.** `var s *MemStore; return s` returns an interface that
  is **not** `nil` — it carries the type `*MemStore` with a nil value. Return a
  bare `nil` instead.
- **Panicking on unknown keys.** The contract here is a nil result, not a panic.
