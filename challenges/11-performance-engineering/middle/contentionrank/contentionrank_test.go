package contentionrank

import (
	"reflect"
	"testing"
)

func TestRankAggregates(t *testing.T) {
	got := Rank([]Record{{"a", 2, 10}, {"a", 2, 30}})
	want := []Site{{"a", 4, 40, 10}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Rank = %v, want %v", got, want)
	}
}

func TestRankOrdersByDelay(t *testing.T) {
	got := Rank([]Record{
		{"small", 100, 100},
		{"big", 1, 5000},
		{"medium", 10, 1000},
	})
	names := []string{got[0].Site, got[1].Site, got[2].Site}
	if !reflect.DeepEqual(names, []string{"big", "medium", "small"}) {
		t.Errorf("order = %v, want [big medium small]", names)
	}
}

func TestRankMeanDelayFindsTheSlowLock(t *testing.T) {
	got := Rank([]Record{
		{"hot-but-quick", 10000, 10000},
		{"rare-but-brutal", 2, 9000},
	})
	if got[0].Site != "hot-but-quick" {
		t.Errorf("first row = %q, want hot-but-quick by total delay", got[0].Site)
	}
	if got[0].MeanDelay != 1 {
		t.Errorf("MeanDelay = %v, want 1", got[0].MeanDelay)
	}
	if got[1].MeanDelay != 4500 {
		t.Errorf("MeanDelay = %v, want 4500", got[1].MeanDelay)
	}
}

func TestRankTieBrokenBySite(t *testing.T) {
	got := Rank([]Record{{"z", 1, 10}, {"a", 1, 10}, {"m", 1, 10}})
	names := []string{got[0].Site, got[1].Site, got[2].Site}
	if !reflect.DeepEqual(names, []string{"a", "m", "z"}) {
		t.Errorf("order = %v, want [a m z]", names)
	}
}

func TestRankDropsJunk(t *testing.T) {
	got := Rank([]Record{{"a", 0, 10}, {"b", -1, 10}, {"c", 1, -5}})
	if got == nil || len(got) != 0 {
		t.Errorf("Rank = %v, want empty non-nil slice", got)
	}
}

func TestRankIsDeterministic(t *testing.T) {
	records := []Record{{"a", 1, 5}, {"b", 1, 5}, {"c", 1, 5}, {"d", 1, 5}}
	first := Rank(records)
	for i := 0; i < 20; i++ {
		if !reflect.DeepEqual(Rank(records), first) {
			t.Fatal("Rank is not deterministic")
		}
	}
}

func TestWorst(t *testing.T) {
	got, ok := Worst([]Record{{"a", 1, 5}, {"b", 1, 50}})
	if !ok || got.Site != "b" || got.Delay != 50 || got.MeanDelay != 50 {
		t.Errorf("Worst = %+v, %v, want site b with delay 50", got, ok)
	}
	if _, ok := Worst(nil); ok {
		t.Error("Worst(nil) reported a site")
	}
	if _, ok := Worst([]Record{{"a", 0, 5}}); ok {
		t.Error("Worst reported a site with no blocking events")
	}
}
