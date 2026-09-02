# Type Switch Dispatch

## Intuition

A type switch is a chain of type-descriptor comparisons resolved at compile time; a reflect-keyed map is a hash lookup resolved at run time. The switch is faster and closed; the table is slower and open to types the package has never heard of.

## Approach

1. `DecodeSwitch` is a plain `switch v.(type)` with a `default`.
2. `Register` keys the label by `reflect.TypeOf(sample)`, ignoring nil.
3. `DecodeTable` returns `"unknown"` for a nil type or an unregistered one.
4. Both paths return identical labels, which the parity test enforces.

## Solution

```go
func DecodeSwitch(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func (t *Table) Register(sample any, label string) {
	rt := reflect.TypeOf(sample)
	if rt == nil {
		return
	}
	t.labels[rt] = label
}

func (t *Table) DecodeTable(v any) string {
	rt := reflect.TypeOf(v)
	if rt == nil {
		return "unknown"
	}
	if label, ok := t.labels[rt]; ok {
		return label
	}
	return "unknown"
}
```

## Walkthrough

`TestNamedTypesAreDistinct` shows both mechanisms agree on Go's rule: `MyInt` is not `int`. The table can then learn `MyInt` at run time — something the switch could only gain by editing its source.

## Pitfalls

- `reflect.TypeOf(nil)` is nil and cannot key a map — guard it in both methods.
- Assuming a type switch and a reflect table differ on named types; both compare exact dynamic types.
- Reaching for reflection when a switch suffices: it is slower and moves errors from compile time to run time.
