# String Pipelines

## Intuition

Chaining methods on a pointer receiver creates a clean DSL for transformations.

## Approach

1. Mutate `p.text` using `strings` functions.
2. Return `p`.

## Solution

```go
func (p *Pipe) Upper() *Pipe {
	p.text = strings.ToUpper(p.text)
	return p
}

func (p *Pipe) Replace(old, new string) *Pipe {
	p.text = strings.ReplaceAll(p.text, old, new)
	return p
}
```

## Walkthrough

- `NewPipe("go lang")` → `text: "go lang"`.
- `.Upper()` → `text: "GO LANG"`, returns pointer.
- `.Replace(" ", "-")` → `text: "GO-LANG"`, returns pointer.

## Pitfalls

- Forgetting to return `p`.
- Forgetting to assign the result of `strings.ToUpper` back to `p.text`.
