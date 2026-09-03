# Per-Tenant API Quota

## Intuition

Two problems stack here: creating a tenant's counter exactly once, and bumping it safely afterwards. `sync.Map.LoadOrStore` solves the first; an `atomic.Int64` inside the stored pointer solves the second.

## Approach

1. `Add`: `LoadOrStore(tenant, &counter{})`, then `.n.Add(n)` on whatever came back.
2. `Used`: `Load`; return 0 when the key is absent.
3. `Tenants`: collect keys with `Range`, then `sort.Strings`.

## Solution

```go
import (
	"sort"
	"sync"
	"sync/atomic"
)

type counter struct{ n atomic.Int64 }

func (q *Quota) Add(tenant string, n int64) int64 {
	v, _ := q.used.LoadOrStore(tenant, &counter{})
	return v.(*counter).n.Add(n)
}

func (q *Quota) Used(tenant string) int64 {
	v, ok := q.used.Load(tenant)
	if !ok {
		return 0
	}
	return v.(*counter).n.Load()
}

func (q *Quota) Tenants() []string {
	var out []string
	q.used.Range(func(k, _ any) bool {
		out = append(out, k.(string))
		return true
	})
	sort.Strings(out)
	return out
}
```

## Walkthrough

Two goroutines charge `acme` at once. Both call `LoadOrStore` with a fresh counter; the map keeps exactly one of them and returns it to both. Each then calls `n.Add(2)` on the same counter, so the total lands at 4 with no lost update.

## Pitfalls

- Storing an `int64` value and doing load-modify-store loses updates.
- `Range` gives no ordering — sort before returning if the caller needs a stable list.
- `sync.Map` is not a general replacement for a mutex-guarded map; it is a fit here because keys are written once.
