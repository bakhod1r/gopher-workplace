package formatter

import "testing"

func TestFormat(t *testing.T) {
	if got := (Plain{}).Format("hi"); got != "hi" {
		t.Errorf("Plain.Format = %q, want \"hi\"", got)
	}
	if got := (KeyValue{}).Format("hi"); got != "msg=hi" {
		t.Errorf("KeyValue.Format = %q, want \"msg=hi\"", got)
	}
}

func TestRender(t *testing.T) {
	cases := []struct {
		name string
		f    Formatter
		msg  string
		want string
	}{
		{"plain", Plain{}, "boom", "boom"},
		{"kv", KeyValue{}, "boom", "msg=boom"},
		{"plain_empty", Plain{}, "", ""},
		{"kv_empty", KeyValue{}, "", "msg="},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Render(tc.f, tc.msg); got != tc.want {
				t.Errorf("Render = %q, want %q", got, tc.want)
			}
		})
	}
}
