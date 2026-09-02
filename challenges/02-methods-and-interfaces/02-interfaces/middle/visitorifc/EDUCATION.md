# Visitor

## Intuition

Two interfaces cross: the node knows how to traverse itself, the visitor knows what to do. Adding an operation costs one new visitor and touches no node type.

## Approach

1. `Text.Accept` visits itself.
2. `Section.Accept` visits itself first (pre-order), then recurses into each child.
3. Each visitor asserts to the node type it cares about and ignores the rest.
4. `Walk` just calls `n.Accept(v)`.

## Solution

```go
func (t Text) Accept(v Visitor) { v.Visit(t) }

func (s Section) Accept(v Visitor) {
	v.Visit(s)
	for _, c := range s.Children {
		c.Accept(v)
	}
}

func (w *WordCounter) Visit(n Node) {
	if t, ok := n.(Text); ok {
		w.Words += len(strings.Fields(t.Content))
	}
}

func (h *HeadingCollector) Visit(n Node) {
	if s, ok := n.(Section); ok {
		h.Titles = append(h.Titles, s.Title)
	}
}

func Walk(n Node, v Visitor) { n.Accept(v) }
```

## Walkthrough

On the test tree the order is: outer section, `"a b"` (2 words), inner section, `"c d e"` (3 words) — 5 words, and headings `[outer inner]` in pre-order.

## Pitfalls

- Visiting children before the section, which reverses the heading order.
- Splitting on `" "` instead of `strings.Fields`, so an empty string counts as one word.
- Asserting to `*Text` when the nodes are stored as values.
