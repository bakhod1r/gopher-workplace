package comparator

import "testing"

func sample() []Record {
	return []Record{{"bob", 30}, {"ann", 25}, {"cid", 35}}
}

func names(recs []Record) []string {
	out := make([]string, len(recs))
	for i, r := range recs {
		out[i] = r.Name
	}
	return out
}

func eq(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestByAge(t *testing.T) {
	if got := names(SortWith(sample(), ByAge{})); !eq(got, []string{"ann", "bob", "cid"}) {
		t.Errorf("ByAge = %v", got)
	}
}

func TestByName(t *testing.T) {
	if got := names(SortWith(sample(), ByName{})); !eq(got, []string{"ann", "bob", "cid"}) {
		t.Errorf("ByName = %v", got)
	}
}

func TestReverse(t *testing.T) {
	if got := names(SortWith(sample(), Reverse{Inner: ByAge{}})); !eq(got, []string{"cid", "bob", "ann"}) {
		t.Errorf("Reverse ByAge = %v", got)
	}
	double := Reverse{Inner: Reverse{Inner: ByAge{}}}
	if got := names(SortWith(sample(), double)); !eq(got, []string{"ann", "bob", "cid"}) {
		t.Errorf("double Reverse = %v", got)
	}
}

func TestInputUnchanged(t *testing.T) {
	in := sample()
	SortWith(in, ByAge{})
	if in[0].Name != "bob" {
		t.Errorf("input was modified: %v", names(in))
	}
}

func TestEmpty(t *testing.T) {
	if got := SortWith(nil, ByAge{}); len(got) != 0 {
		t.Errorf("SortWith(nil) = %v", got)
	}
}
