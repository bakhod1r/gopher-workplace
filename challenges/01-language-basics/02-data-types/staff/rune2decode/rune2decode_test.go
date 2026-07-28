package rune2decode

import "testing"

func TestDecode2(t *testing.T) {
	cases := []struct {
		lead, cont byte
		want       rune
	}{
		{0xC3, 0xA9, 'é'}, // U+00E9
		{0xC2, 0xA9, '©'}, // U+00A9
		{0xD0, 0x81, 'Ё'}, // U+0401 (lead bit needs 0x1F mask)
		{0xD8, 0xA7, 'ا'}, // U+0627 Arabic (exercises the 5th payload bit)
	}
	for _, c := range cases {
		if got := Decode2(c.lead, c.cont); got != c.want {
			t.Errorf("Decode2(%#x,%#x)=%q(%U); want %q(%U)", c.lead, c.cont, got, got, c.want, c.want)
		}
	}
}
