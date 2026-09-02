package uploadfeed

import "testing"

func TestUploadFeed(t *testing.T) {
	cases := []struct {
		name string
		keys []string
		want []string
	}{
		{"batch_of_three", []string{"a.jpg", "b.jpg", "c.jpg"}, []string{"a.jpg", "b.jpg", "c.jpg"}},
		{"single_object", []string{"only.png"}, []string{"only.png"}},
		{"duplicate_keys", []string{"x.gif", "x.gif"}, []string{"x.gif", "x.gif"}},
		{"empty_batch", []string{}, nil},
		{"nil_batch", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for k := range UploadFeed(tc.keys) {
				got = append(got, k)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("UploadFeed(%v) = %v, want %v", tc.keys, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("UploadFeed(%v) = %v, want %v", tc.keys, got, tc.want)
				}
			}
		})
	}
}
