package csvline

import "testing"

func chanOf(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestJoinFields(t *testing.T) {
	cases := []struct {
		name   string
		fields []string
		sep    string
		want   string
	}{
		{"two_fields", []string{"a", "b"}, ",", "a,b"},
		{"single_field", []string{"x"}, ",", "x"},
		{"empty_record", nil, ",", ""},
		{"tab_delimited", []string{"a", "b", "c"}, "\t", "a\tb\tc"},
		{"empty_sep", []string{"go", "lang"}, "", "golang"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := JoinFields(chanOf(tc.fields...), tc.sep); got != tc.want {
				t.Errorf("JoinFields(%v, %q) = %q, want %q",
					tc.fields, tc.sep, got, tc.want)
			}
		})
	}
}
