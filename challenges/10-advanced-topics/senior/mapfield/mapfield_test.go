package mapfield

import (
	"errors"
	"testing"
)

type doc struct {
	Tags map[string]string
	Name string
}

func TestPutTagCreatesTheMap(t *testing.T) {
	d := &doc{}
	if err := PutTag(d, "a", "1"); err != nil {
		t.Fatal(err)
	}
	if d.Tags == nil {
		t.Fatal("the map was not created")
	}
	if d.Tags["a"] != "1" {
		t.Errorf("Tags = %v, want map[a:1]", d.Tags)
	}
}

func TestPutTagReusesTheMap(t *testing.T) {
	d := &doc{Tags: map[string]string{"keep": "yes"}}
	existing := d.Tags
	if err := PutTag(d, "a", "1"); err != nil {
		t.Fatal(err)
	}
	if d.Tags["keep"] != "yes" || d.Tags["a"] != "1" {
		t.Errorf("Tags = %v, want both entries", d.Tags)
	}
	d.Tags["direct"] = "x"
	if existing["direct"] != "x" {
		t.Error("the field was replaced with a new map")
	}
}

func TestPutTagOverwrites(t *testing.T) {
	d := &doc{}
	PutTag(d, "a", "1")
	if err := PutTag(d, "a", "2"); err != nil {
		t.Fatal(err)
	}
	if d.Tags["a"] != "2" {
		t.Errorf("Tags[a] = %q, want \"2\"", d.Tags["a"])
	}
}

func TestPutTagBadTargets(t *testing.T) {
	type noTags struct{ A int }
	type wrongKind struct{ Tags []string }
	type wrongTypes struct{ Tags map[string]int }
	type unexported struct{ tags map[string]string }

	cases := []any{
		doc{}, nil, (*doc)(nil), new(int),
		&noTags{}, &wrongKind{}, &wrongTypes{}, &unexported{},
	}
	for _, c := range cases {
		if err := PutTag(c, "a", "1"); !errors.Is(err, ErrTarget) {
			t.Errorf("PutTag(%#v) = %v, want ErrTarget", c, err)
		}
	}
}
