# Type Switch

## Intuition

An interface value carries its dynamic type. A type switch reads that type once and gives you a correctly typed variable per branch — cheaper and clearer than a chain of assertions.

## Approach

1. Write `switch x := v.(type)`.
2. Handle `int` with `strconv.Itoa`, `bool` with an explicit `if`, `string` directly.
3. Handle `[]string` with `strings.Join(x, ",")`.
4. Return `"?"` from `default`, which also catches `nil`.

## Solution

```go
func Render(v any) string {
	switch x := v.(type) {
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	case string:
		return x
	case []string:
		return strings.Join(x, ",")
	default:
		return "?"
	}
}
```

## Walkthrough

`Render(nil)`: a nil `any` has no dynamic type, so no case matches and `default` returns `"?"`. `Render([]int{1})` has a dynamic type that is not listed, so it also falls to `default`.

## Pitfalls

- Adding `case nil:` is legal but unnecessary here — `default` already covers it.
- `fmt.Sprint(x)` for every case: it works for ints but prints `[a b]` for a slice.
- Assuming `int32` or `int64` match `case int` — they do not; the dynamic type must be exact.
