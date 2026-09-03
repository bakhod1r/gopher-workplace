# Region Replication Tree

## Intuition

A goroutine tree is only safe if every parent joins its own children. Get that right and the single outer `Wait` in the coordinator is a complete guarantee about the whole tree, however deep it goes. Get it wrong and a zone goroutine outlives the call that spawned it, writing into a slice the caller already returned.

## Approach

1. Declare `mu`, `failed []string`, and the outer WaitGroup in the coordinator.
2. Launch a goroutine per region, taking the `Region` value as a parameter.
3. Inside, declare a fresh inner WaitGroup, launch a goroutine per zone, and `inner.Wait()` before returning.
4. In a zone goroutine, on error lock the mutex and append `region.Name + "/" + zone`.
5. After `outer.Wait()`, sort and normalise nil to `[]string{}`.

## Solution

```go
// Package regionreplicator — Gopher Workplace challenge.
package regionreplicator

import (
	"sort"
	"sync"
)

// Region is one geographic region and the availability zones inside it.
type Region struct {
	Name  string
	Zones []string
}

// ReplicateAll fans out one goroutine per region, and inside each region one
// goroutine per zone — a two-level tree. It returns the "region/zone" pairs that
// failed to replicate, sorted, so the runbook lists the same targets every time.
//
// Examples:
//
//	ReplicateAll([]Region{{"eu", []string{"a"}}}, replicate)          => []
//	ReplicateAll([]Region{{"eu", []string{"a", "bad"}}}, replicate)   => ["eu/bad"]
//	ReplicateAll(nil, replicate)                                      => []
func ReplicateAll(regions []Region, replicate func(region, zone string) error) []string {
	var (
		mu     sync.Mutex
		failed []string
		outer  sync.WaitGroup
	)

	for _, region := range regions {
		outer.Add(1)
		go func(region Region) {
			defer outer.Done()

			var inner sync.WaitGroup
			for _, zone := range region.Zones {
				inner.Add(1)
				go func(zone string) {
					defer inner.Done()
					if err := replicate(region.Name, zone); err != nil {
						mu.Lock()
						failed = append(failed, region.Name+"/"+zone)
						mu.Unlock()
					}
				}(zone)
			}
			inner.Wait()
		}(region)
	}
	outer.Wait()

	sort.Strings(failed)
	if failed == nil {
		return []string{}
	}
	return failed
}
```

## Walkthrough

- For two regions with three zones total, three leaf goroutines run and every one calls `replicate` exactly once.
- `failures_sorted_across_regions` produces `us/bad-1` and `ap/bad-2` from two different subtrees; the sort yields `[ap/bad-2 us/bad-1]`.
- `whole_region_down` shows both children of a single region reporting independently.
- A region with no zones spawns no leaves; its inner `Wait` returns immediately and the outer group still balances.

## Pitfalls

- Hoisting the inner WaitGroup out of the region goroutine, coupling unrelated regions together.
- Skipping `inner.Wait()`, so the region reports done while its zones are still appending — a use-after-return on the result slice.
- Capturing the loop variables instead of passing `region` and `zone` in as parameters.
- Appending without the mutex; the race detector fails the build.
