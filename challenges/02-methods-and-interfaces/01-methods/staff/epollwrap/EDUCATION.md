# Epoll Interest Set

## Intuition

`epoll_wait` hands back a readiness report, not a to-do list. Two filters must
be applied before it means anything: *is this descriptor still mine?* and *is
the event that fired one I asked for?* Both are cheap — a map lookup and a
bitwise AND — and both are easy to forget, which is how event loops end up
spinning on descriptors nobody is watching.

## Approach

1. Iterate the readiness report.
2. Look the descriptor up in the interest map with the comma-ok form; skip
   unknown descriptors.
3. Keep the descriptor when the masks share at least one bit.
4. Sort, because map iteration order is deliberately randomized.

## Solution

```go
func (e *Epoll) Wait(ready map[int]uint32) []int {
	var fds []int
	for fd, events := range ready {
		interest, ok := e.interest[fd]
		if !ok {
			continue
		}
		if events&interest != 0 {
			fds = append(fds, fd)
		}
	}
	sort.Ints(fds)
	return fds
}
```

## Walkthrough

With fd 3 registered for read, 7 for write and 5 for both:

| fd | ready | interest | `&` | kept |
|----|-------|----------|-----|------|
| 3 | read | read | read | yes |
| 7 | read | write | 0 | no |
| 5 | write | read\|write | write | yes |
| 9 | read\|write | — | unregistered | no |

The survivors are 3 and 5. They may have been visited in either order, so
`sort.Ints` is what makes `reflect.DeepEqual` against `[]int{3, 5}` meaningful.

## Pitfalls

- **`e.interest[fd] != 0` instead of comma-ok.** It happens to behave the same
  while every registration has a non-zero mask, but a descriptor registered with
  mask 0 is legal, and the two cases mean different things.
- **`events == interest`.** Requires an exact match, so fd 5 — interested in two
  events, one of which fired — is dropped.
- **Skipping the sort.** The test passes perhaps a third of the time; Go
  randomizes map iteration order precisely to expose this class of bug early.
- **Returning `[]int{}` vs `nil`.** Both have length 0 and both satisfy the
  tests here, which check `len`.

## Why map order is random

Go's runtime picks a random start bucket for every `range` over a map. It is a
deliberate anti-feature: it makes accidental order dependence fail fast in
development instead of after a map grows in production.
