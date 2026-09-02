package typeswitchgen

import "testing"

type Celsius float64

func TestDescribe(t *testing.T) {
	if got := Describe(1); got != "integer" {
		t.Errorf("Describe(1) = %q, want integer", got)
	}
	if got := Describe(int64(1)); got != "integer" {
		t.Errorf("Describe(int64) = %q, want integer", got)
	}
	if got := Describe(1.5); got != "float" {
		t.Errorf("Describe(1.5) = %q, want float", got)
	}
	if got := Describe("a"); got != "string" {
		t.Errorf(`Describe("a") = %q, want string`, got)
	}
	if got := Describe(true); got != "other" {
		t.Errorf("Describe(true) = %q, want other", got)
	}
}

func TestDescribeNamedType(t *testing.T) {
	if got := Describe(Celsius(20)); got != "other" {
		t.Errorf("Describe(Celsius) = %q, want other (the dynamic type is exact)", got)
	}
}
