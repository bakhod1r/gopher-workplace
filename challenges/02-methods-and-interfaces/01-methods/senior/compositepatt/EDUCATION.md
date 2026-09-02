# Composite Pattern

## Intuition

The point of the composite is that a leaf and a container are asked the same
question. `Size()` does not care whether it is standing on a folder full of
files or a folder full of folders — it adds up what it can see and delegates
the rest to the same method one level down.

## Approach

1. Start a total at 0.
2. Add every file size directly.
3. For each subfolder, add the answer *it* gives.

## Solution

```go
func (f *Folder) Size() int {
	total := 0
	for _, size := range f.Files {
		total += size
	}
	for _, sub := range f.Sub {
		total += sub.Size()
	}
	return total
}
```

## Walkthrough

For the test tree, the root loop adds `10 + 20 = 30`. Then `Sub[0].Size()` runs
the same code with `Files = [30]` and an empty `Sub`, so its second loop never
executes and it returns `30`. `Sub[1].Size()` returns `40 + 50 = 90`. Root
total: `30 + 30 + 90 = 150`.

## Pitfalls

- **Summing `len(f.Files)` instead of the values.** That counts files, not bytes.
- **Recursing into `f` itself.** `total += f.Size()` is infinite recursion and a
  stack overflow.
- **Forgetting one of the two loops.** The test tree has weight at both levels
  precisely to catch that.

## Why a method, not a function

`Size` could be `func Size(f *Folder) int`. Making it a method means any type
that can report a size satisfies the same one-method interface later — which is
how the composite grows into a real tree of mixed node types.
