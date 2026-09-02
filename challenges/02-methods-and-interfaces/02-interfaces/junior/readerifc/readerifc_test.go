package readerifc

import "testing"

func TestRead(t *testing.T) {
	s := &StringSource{Data: "abc", Chunk: 2}

	if c, ok := s.Read(); c != "ab" || !ok {
		t.Errorf("Read 1 = %q, %v, want \"ab\", true", c, ok)
	}
	if c, ok := s.Read(); c != "c" || !ok {
		t.Errorf("Read 2 = %q, %v, want \"c\", true", c, ok)
	}
	if c, ok := s.Read(); c != "" || ok {
		t.Errorf("Read 3 = %q, %v, want \"\", false", c, ok)
	}
}

func TestReadAll(t *testing.T) {
	cases := []struct {
		name  string
		data  string
		chunk int
		want  string
	}{
		{"even", "hello", 5, "hello"},
		{"uneven", "hello", 2, "hello"},
		{"one_byte", "go", 1, "go"},
		{"empty", "", 3, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReadAll(&StringSource{Data: tc.data, Chunk: tc.chunk})
			if got != tc.want {
				t.Errorf("ReadAll = %q, want %q", got, tc.want)
			}
		})
	}
}
