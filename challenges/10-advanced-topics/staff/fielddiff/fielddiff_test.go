package fielddiff

import (
	"reflect"
	"testing"
)

type limits struct {
	Soft int
	Hard int
}

type settings struct {
	Name   string
	Retry  int
	Limits limits
	hidden int
}

func TestDiffFlat(t *testing.T) {
	a := settings{Name: "x", Retry: 1}
	b := settings{Name: "y", Retry: 1}
	if got := Diff(a, b); !reflect.DeepEqual(got, []string{"Name"}) {
		t.Errorf("Diff = %v, want [Name]", got)
	}
}

func TestDiffNested(t *testing.T) {
	a := settings{Limits: limits{Soft: 1, Hard: 2}}
	b := settings{Limits: limits{Soft: 9, Hard: 2}}
	if got := Diff(a, b); !reflect.DeepEqual(got, []string{"Limits.Soft"}) {
		t.Errorf("Diff = %v, want [Limits.Soft]", got)
	}
}

func TestDiffMultipleInDeclarationOrder(t *testing.T) {
	a := settings{Name: "x", Retry: 1, Limits: limits{Soft: 1}}
	b := settings{Name: "y", Retry: 2, Limits: limits{Soft: 2}}
	want := []string{"Name", "Retry", "Limits.Soft"}
	if got := Diff(a, b); !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffIdentical(t *testing.T) {
	a := settings{Name: "x", Retry: 1, Limits: limits{Soft: 1, Hard: 2}}
	if got := Diff(a, a); len(got) != 0 {
		t.Errorf("Diff = %v, want empty", got)
	}
}

func TestDiffIgnoresUnexported(t *testing.T) {
	a := settings{hidden: 1}
	b := settings{hidden: 2}
	if got := Diff(a, b); len(got) != 0 {
		t.Errorf("Diff = %v, want empty: unexported fields are not compared", got)
	}
}

func TestDiffMismatchedTypes(t *testing.T) {
	if got := Diff(settings{}, limits{}); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
	if got := Diff(nil, settings{}); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
	if got := Diff(nil, nil); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
}

func TestDiffScalars(t *testing.T) {
	if got := Diff(1, 2); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("Diff = %v, want [\"\"]", got)
	}
	if got := Diff(1, 1); len(got) != 0 {
		t.Errorf("Diff = %v, want empty", got)
	}
}
