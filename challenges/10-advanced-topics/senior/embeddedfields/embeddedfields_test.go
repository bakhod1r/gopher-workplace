package embeddedfields

import (
	"reflect"
	"testing"
)

type Base struct {
	ID     int
	hidden int
}

type Meta struct {
	Tag string
}

type User struct {
	Base
	Name  string
	Extra Meta
}

type Deep struct {
	User
	Note string
}

func TestPaths(t *testing.T) {
	want := []string{"Base.ID", "Name", "Extra"}
	if got := Paths(User{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Paths = %v, want %v", got, want)
	}
}

func TestPathsNested(t *testing.T) {
	want := []string{"User.Base.ID", "User.Name", "User.Extra", "Note"}
	if got := Paths(Deep{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Paths = %v, want %v", got, want)
	}
}

func TestPathsSkipsUnexported(t *testing.T) {
	got := Paths(Base{})
	if !reflect.DeepEqual(got, []string{"ID"}) {
		t.Errorf("Paths = %v, want [ID]", got)
	}
}

func TestPathsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &User{}} {
		if got := Paths(in); got != nil {
			t.Errorf("Paths(%#v) = %v, want nil", in, got)
		}
	}
}

func TestPathsNamedStructFieldIsALeaf(t *testing.T) {
	got := Paths(User{})
	for _, p := range got {
		if p == "Extra.Tag" {
			t.Error("a named struct field must not be descended into")
		}
	}
}
