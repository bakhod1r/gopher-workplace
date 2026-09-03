package routetable

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestLookup(t *testing.T) {
	cases := []struct {
		name    string
		set     [][2]string
		path    string
		want    string
		wantHit bool
	}{
		{"hit", [][2]string{{"/api", "backend-1"}}, "/api", "backend-1", true},
		{"miss", [][2]string{{"/api", "backend-1"}}, "/other", "", false},
		{"empty_table", nil, "/api", "", false},
		{"overwrite", [][2]string{{"/api", "backend-1"}, {"/api", "backend-2"}}, "/api", "backend-2", true},
		{"two_routes", [][2]string{{"/a", "b1"}, {"/b", "b2"}}, "/b", "b2", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tbl := NewTable()
			for _, kv := range tc.set {
				tbl.Set(kv[0], kv[1])
			}
			got, hit := tbl.Lookup(tc.path)
			if got != tc.want || hit != tc.wantHit {
				t.Errorf("Lookup(%q) = %q, %v; want %q, %v", tc.path, got, hit, tc.want, tc.wantHit)
			}
		})
	}
}

func TestSnapshotIsACopy(t *testing.T) {
	tbl := NewTable()
	tbl.Set("/api", "backend-1")

	snap := tbl.Snapshot()
	want := map[string]string{"/api": "backend-1"}
	if !reflect.DeepEqual(snap, want) {
		t.Fatalf("Snapshot() = %v, want %v", snap, want)
	}

	snap["/api"] = "tampered"
	if got, _ := tbl.Lookup("/api"); got != "backend-1" {
		t.Errorf("mutating the snapshot changed the table: %q", got)
	}

	tbl.Set("/new", "backend-2")
	if _, ok := snap["/new"]; ok {
		t.Error("snapshot saw a route added after it was taken")
	}
}

func TestConcurrentReadersAndWriter(t *testing.T) {
	tbl := NewTable()
	const routes = 50
	for i := range routes {
		tbl.Set("/p"+strconv.Itoa(i), "backend-"+strconv.Itoa(i))
	}

	var wg sync.WaitGroup
	for i := range routes {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			path := "/p" + strconv.Itoa(i)
			if got, ok := tbl.Lookup(path); !ok || got != "backend-"+strconv.Itoa(i) {
				t.Errorf("Lookup(%q) = %q, %v", path, got, ok)
			}
			tbl.Snapshot()
		}(i)
	}
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tbl.Set("/deploy"+strconv.Itoa(i), "canary")
		}(i)
	}
	wg.Wait()

	if got := tbl.Len(); got != routes+10 {
		t.Errorf("Len() = %d, want %d", got, routes+10)
	}
}
