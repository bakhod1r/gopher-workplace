package ifacecomp

import "testing"

func TestFileReadWrite(t *testing.T) {
	f := &File{}
	f.Write("hello")
	if got := f.Read(); got != "hello" {
		t.Errorf("Read = %q, want \"hello\"", got)
	}

	var rw ReadWriter = f
	rw.Write("again")
	if got := rw.Read(); got != "again" {
		t.Errorf("Read = %q, want \"again\"", got)
	}
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		name string
		v    any
		want string
	}{
		{"file", &File{}, "rw"},
		{"read_only", ReadOnly{Data: "x"}, "r"},
		{"write_only", &WriteOnly{}, "w"},
		{"neither", struct{}{}, "none"},
		{"file_value", File{}, "none"},
		{"nil", nil, "none"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.v); got != tc.want {
				t.Errorf("Describe(%T) = %q, want %q", tc.v, got, tc.want)
			}
		})
	}
}
