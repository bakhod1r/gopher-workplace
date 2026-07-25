package checksum

import "testing"

func TestChecksum(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want uint8
	}{
		{"nil", nil, 0},
		{"empty", []byte{}, 0},
		{"no wrap", []byte{1, 2, 3}, 6},
		{"exact wrap to zero", []byte{255, 1}, 0},
		{"wrap 300", []byte{200, 100}, 44},
		{"single max byte", []byte{255}, 255},
		{"many wraps", []byte{255, 255, 255, 255}, 252}, // 1020 mod 256
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Checksum(tc.in); got != tc.want {
				t.Errorf("Checksum(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
