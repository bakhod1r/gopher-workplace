# Method Sets

## Intuition

The method set decides interface satisfaction. `*T` includes methods declared on both `T` and `*T`; `T` includes only those on `T`. Storing a value in an interface copies it, and a copy has no address, so pointer methods cannot apply.

## Approach

1. Give `Value` a value-receiver `Name`, and `Pointer` pointer-receiver `Name` and `Rename`.
2. `Names` calls `n.Name()` for each element.
3. `Satisfies` uses a comma-ok assertion to `Renamer` and returns only the bool.

## Solution

```go
func (v Value) Name() string { return v.N }

func (p *Pointer) Name() string { return p.N }

func (p *Pointer) Rename(s string) { p.N = s }

func Names(ns []Named) []string {
	out := make([]string, 0, len(ns))
	for _, n := range ns {
		out = append(out, n.Name())
	}
	return out
}

func Satisfies(v any) bool {
	_, ok := v.(Renamer)
	return ok
}
```

## Walkthrough

`Satisfies(&Value{})` is false: `Value` has no `Rename` at all, so neither `Value` nor `*Value` satisfies `Renamer`. `Satisfies(Pointer{})` is false because `Rename` lives on `*Pointer`.

## Pitfalls

- Expecting `Pointer{}` to satisfy `Renamer` because `p.Rename(...)` compiles on a local variable — that is addressability sugar, not method-set membership.
- Giving `Pointer` a value-receiver `Name` and a pointer-receiver `Rename`, which makes the method set inconsistent.
- Assuming `&Value{}` gains methods it never declared.
