package zerocopy

import "testing"

func strs(fields [][]byte) []string {
	out := make([]string, len(fields))
	for i, f := range fields {
		out[i] = string(f)
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

func TestFields(t *testing.T) {
	p := &ZeroCopyParser{Sep: ','}
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"three", "a,bb,c", []string{"a", "bb", "c"}},
		{"one", "solo", []string{"solo"}},
		{"trailing_sep", "a,", []string{"a", ""}},
		{"leading_sep", ",a", []string{"", "a"}},
		{"only_seps", ",,", []string{"", "", ""}},
		{"empty", "", nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := strs(p.Fields([]byte(tc.in)))
			if !eq(got, tc.want) {
				t.Errorf("Fields = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFieldsAliasInput(t *testing.T) {
	p := &ZeroCopyParser{Sep: ','}
	data := []byte("ab,cd")
	fields := p.Fields(data)

	data[0] = 'X'
	if string(fields[0]) != "Xb" {
		t.Errorf("field = %q; Fields must alias the input, not copy it", fields[0])
	}
}

func TestCopyFieldsAreIndependent(t *testing.T) {
	p := &ZeroCopyParser{Sep: ','}
	data := []byte("ab,cd")
	fields := CopyFields(p, data)

	data[0] = 'X'
	if string(fields[0]) != "ab" {
		t.Errorf("field = %q; CopyFields must return independent copies", fields[0])
	}
}

func TestNoPerFieldAllocation(t *testing.T) {
	p := &ZeroCopyParser{Sep: ','}
	data := []byte("aaaa,bbbb,cccc,dddd,eeee,ffff,gggg,hhhh")

	avg := testing.AllocsPerRun(200, func() {
		_ = p.Fields(data)
	})
	if avg > 1 {
		t.Errorf("Fields allocated %.2f times per call; only the result slice may allocate", avg)
	}
}

func BenchmarkFields(b *testing.B) {
	p := &ZeroCopyParser{Sep: ','}
	data := []byte("aaaa,bbbb,cccc,dddd,eeee,ffff,gggg,hhhh")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = p.Fields(data)
	}
}

func BenchmarkCopyFields(b *testing.B) {
	p := &ZeroCopyParser{Sep: ','}
	data := []byte("aaaa,bbbb,cccc,dddd,eeee,ffff,gggg,hhhh")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = CopyFields(p, data)
	}
}
