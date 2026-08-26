# Methods on String Types

## Intuition

`type MyString string` creates a new type. You can add methods to `MyString`
but not to `string`. This is how Go extends built-in types without inheritance.

## Approach

1. Convert receiver to `string`.
2. Call `strings.ToUpper`.
3. Return the result.

## Solution

```go
func (s MyString) Upper() string {
	return strings.ToUpper(string(s))
}
```

## Walkthrough

For `MyString("hello")`:
- `string(s)` → `"hello"`.
- `strings.ToUpper("hello")` → `"HELLO"`.

## Pitfalls

- Calling `strings.ToUpper(s)` without conversion — `MyString` is not `string`.
- Returning `MyString` instead of `string` — the signature says `string`.
