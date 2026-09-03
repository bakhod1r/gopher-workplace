package tenantquota

import (
	"reflect"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name   string
		tenant string
		steps  []int64
		want   int64
	}{
		{"single", "acme", []int64{3}, 3},
		{"accumulates", "acme", []int64{3, 2}, 5},
		{"zero", "acme", []int64{0}, 0},
		{"negative_refund", "acme", []int64{10, -4}, 6},
		{"many", "globex", []int64{1, 1, 1, 1}, 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var q Quota
			var got int64
			for _, n := range tc.steps {
				got = q.Add(tc.tenant, n)
			}
			if got != tc.want {
				t.Errorf("Add total = %d, want %d", got, tc.want)
			}
			if used := q.Used(tc.tenant); used != tc.want {
				t.Errorf("Used(%q) = %d, want %d", tc.tenant, used, tc.want)
			}
		})
	}
}

func TestUsedUnknownTenant(t *testing.T) {
	var q Quota
	if got := q.Used("nobody"); got != 0 {
		t.Errorf("Used(unknown) = %d, want 0", got)
	}
}

func TestTenantsSorted(t *testing.T) {
	var q Quota
	q.Add("b", 1)
	q.Add("a", 1)
	q.Add("c", 1)
	want := []string{"a", "b", "c"}
	if got := q.Tenants(); !reflect.DeepEqual(got, want) {
		t.Errorf("Tenants() = %v, want %v", got, want)
	}
}

func TestAddConcurrent(t *testing.T) {
	var q Quota
	const tenants, perTenant = 4, 250

	var wg sync.WaitGroup
	names := []string{"acme", "globex", "initech", "umbrella"}
	for _, name := range names {
		for range perTenant {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				q.Add(name, 2)
			}(name)
		}
	}
	wg.Wait()

	for _, name := range names {
		if got := q.Used(name); got != perTenant*2 {
			t.Errorf("Used(%q) = %d, want %d", name, got, perTenant*2)
		}
	}
	if got := len(q.Tenants()); got != tenants {
		t.Errorf("len(Tenants()) = %d, want %d", got, tenants)
	}
}
