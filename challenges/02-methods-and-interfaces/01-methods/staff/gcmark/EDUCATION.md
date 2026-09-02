# GC Mark Phase

## Intuition

Marking is graph traversal where the "visited" set costs nothing extra: the mark
bit you are setting anyway *is* the visited set. That is why real collectors can
trace a heap with cycles without a side table.

## Approach

1. Stop on nil and on an already-marked object.
2. Set the mark.
3. Recurse into every reference.

## Solution

```go
func (o *Object) Mark() {
	if o == nil || o.Marked {
		return
	}
	o.Marked = true
	for _, ref := range o.Refs {
		ref.Mark()
	}
}
```

## Walkthrough

`o1.Mark()` sets `o1.Marked` and iterates its one reference, `o2`. `o2` is
unmarked and has no refs, so it marks itself and returns. The test asserts
`o2.Marked`, proving the trace crossed the edge.

Add `o2.Refs = append(o2.Refs, o1)` to make a cycle: the recursion reaches `o1`
again, finds `Marked == true`, and returns immediately.

## Pitfalls

- **Marking after the loop.** The mark bit is then not yet set when the cycle
  comes back around, and the recursion never terminates.
- **Omitting the `o.Marked` guard.** Fine for a tree, fatal for a graph —
  `stack overflow` on the first cycle.
- **Value receiver.** The mark lands on a copy; nothing is ever collected, or
  rather, everything is.

## The other half: sweep

After marking, the collector walks the whole heap and frees anything unmarked,
then clears the marks for the next cycle. The `gcswp` puzzle covers that side.
The pair — mark then sweep — is the oldest tracing algorithm there is, and Go's
own concurrent collector is still recognisably a descendant of it.
