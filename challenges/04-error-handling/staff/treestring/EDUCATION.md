# Render The Tree

## Intuition

Rendering the structure rather than the combined message shows which annotation belongs to which branch — the thing a flat message cannot express.

## Approach

1. Recurse with a depth parameter.
2. Prefix each line with `depth` tabs.
3. Recurse into joined branches or the wrapped child, then join the lines.

## Solution

```go
// Tree:
lines := render(err, 0)
return strings.Join(lines, "\n")

// render:
if err == nil {
	return nil
}
out := []string{strings.Repeat("\t", depth) + err.Error()}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		out = append(out, render(e, depth+1)...)
	}
	return out
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	out = append(out, render(wrapped.Unwrap(), depth+1)...)
}
return out
```

## Walkthrough

A joined error prints its own combined message first — which already contains both branches — then each branch indented one level.

## Pitfalls

- Indenting the root, producing a leading tab.
- Rendering children before the parent.
- Handling only one unwrap shape, so joins print as a single line.
