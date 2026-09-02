# False Sharing

## Intuition

Cache coherence works at cache-line granularity, not variable granularity. Two counters 8 bytes apart live on one line, so every increment on either core invalidates the other's copy — the cores serialise on hardware you cannot see in the source.

## Approach

1. `paddedCell` holds an 8-byte counter plus `CacheLine-8` bytes of padding.
2. Both `Inc` methods bounds-check, then use `atomic.Int64.Add`.
3. Both `Total` methods sum with `Load`.
4. The tests verify the layout with `unsafe.Sizeof` and pointer arithmetic.

## Solution

```go
func (p *PaddedCounters) Inc(i int) {
	if i < 0 || i >= len(p.cells) {
		return
	}
	p.cells[i].v.Add(1)
}

func (p *PaddedCounters) Total() int64 {
	var sum int64
	for i := range p.cells {
		sum += p.cells[i].v.Load()
	}
	return sum
}

func (p *PackedCounters) Inc(i int) {
	if i < 0 || i >= len(p.vals) {
		return
	}
	p.vals[i].Add(1)
}
```

## Walkthrough

Both variants produce identical results; only their throughput differs. Run the two parallel benchmarks on a multi-core machine and the packed version is dramatically slower despite doing exactly the same atomic work.

## Pitfalls

- Padding the struct but storing pointers to cells, which puts the counters back on shared lines.
- Assuming a 64-byte line everywhere — some architectures use 128; measure rather than assume.
- Padding cold data, which wastes memory and cache capacity for no gain.
