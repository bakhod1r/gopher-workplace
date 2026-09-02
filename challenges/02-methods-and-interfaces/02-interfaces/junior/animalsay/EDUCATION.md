# Animal Say

## Intuition

`Dog` and `Cat` never mention `Animal`. They satisfy it simply by having `Sound() string` — satisfaction in Go is structural.

## Approach

1. Return the fixed string from each `Sound`.
2. `MakeNoise` returns `a.Sound()`.
3. `Chorus` appends `" "` before every element after the first.

## Solution

```go
func (d Dog) Sound() string { return "Woof!" }

func (c Cat) Sound() string { return "Meow!" }

func MakeNoise(a Animal) string { return a.Sound() }

func Chorus(as []Animal) string {
	out := ""
	for i, a := range as {
		if i > 0 {
			out += " "
		}
		out += a.Sound()
	}
	return out
}
```

## Walkthrough

`Chorus([]Animal{Dog{}, Cat{}})`: i=0 adds `Woof!`; i=1 adds a space then `Meow!` — `"Woof! Meow!"`.

## Pitfalls

- A trailing space from adding the separator after each sound.
- Dropping the exclamation marks.
- Type-switching in `MakeNoise` instead of calling the method.
