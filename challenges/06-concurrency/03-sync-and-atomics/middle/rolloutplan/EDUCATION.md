# Copy-on-Write Rollout Plan

## Intuition

Readers here are wait-free: `current.Load()` is one instruction and the map it returns will never change again. The whole cost of concurrency moves onto the writer, which is exactly the right trade when writes are thousands of times rarer than reads. The rule that makes it safe is absolute — once a map is stored, nobody may write to it, ever.

## Approach

1. `NewPlan`: allocate a map, copy `initial` through `clamp`, `Store` its address.
2. `Percent`: `(*p.current.Load())[feature]` — a missing key gives the zero value.
3. `SetPercent`: `Lock`; dereference the current pointer; build a new map of `len(old)+1`; copy; set the clamped value; `Store` the new pointer; `Unlock`.
4. `Features`: load the snapshot once into a local, collect its keys, `sort.Strings`.

## Solution

```go
// NewPlan returns a Plan seeded with a copy of initial. Percentages are
// clamped into 0..100.
//
// Examples:
//
//	NewPlan(map[string]int{"checkout": 25}).Percent("checkout") => 25
//	NewPlan(nil).Percent("checkout")                            => 0
func NewPlan(initial map[string]int) *Plan {
	p := &Plan{}
	seed := make(map[string]int, len(initial))
	for name, pct := range initial {
		seed[name] = clamp(pct)
	}
	p.current.Store(&seed)
	return p
}

// Percent returns a feature's rollout percentage, or 0 if it is unknown.
// It takes no lock: it reads whichever snapshot is current.
//
// Examples:
//
//	NewPlan(nil).Percent("unknown") => 0
func (p *Plan) Percent(feature string) int {
	return (*p.current.Load())[feature]
}

// SetPercent publishes a new snapshot with feature set to pct, clamped into
// 0..100. The previous snapshot is left untouched for readers still using it.
//
// Examples:
//
//	p := NewPlan(nil); p.SetPercent("checkout", 40); p.Percent("checkout") => 40
//	p.SetPercent("checkout", 300); p.Percent("checkout")                   => 100
func (p *Plan) SetPercent(feature string, pct int) {
	p.writers.Lock()
	defer p.writers.Unlock()

	old := *p.current.Load()
	next := make(map[string]int, len(old)+1)
	for name, v := range old {
		next[name] = v
	}
	next[feature] = clamp(pct)
	p.current.Store(&next)
}

// Features returns every feature in the current snapshot, sorted.
//
// Examples:
//
//	p.SetPercent("b", 1); p.SetPercent("a", 1); p.Features() => ["a" "b"]
func (p *Plan) Features() []string {
	snap := *p.current.Load()
	names := make([]string, 0, len(snap))
	for name := range snap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
```

## Walkthrough

- A reader that loaded snapshot A keeps reading A safely even while a writer publishes B; A is garbage-collected when the last reader drops it.
- Two concurrent `SetPercent` calls serialise on `p.writers`, so the second one copies the first one's map and neither update is lost.
- `NewPlan` copying `initial` is what makes `TestPlanDoesNotAliasInitial` pass — sharing the caller's map would leak an unsynchronised writer into the plan.
- `Features` loads the pointer once; loading it again mid-loop could mix two snapshots.

## Pitfalls

- Mutating the loaded map in place (`(*p.current.Load())[f] = pct`) — that is a map write racing every reader, and the entire design collapses.
- Dropping the writer mutex: two writers both copy the same old map and one update vanishes.
- Loading `current` twice inside one method and assuming both loads saw the same snapshot.
- Storing `&someLocalMapVariable` that you keep writing to afterwards; publish only maps you are done with.
