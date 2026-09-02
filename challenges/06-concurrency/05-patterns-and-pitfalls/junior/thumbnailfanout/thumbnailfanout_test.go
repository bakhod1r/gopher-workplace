package thumbnailfanout

import "testing"

func TestRenderThumbnails(t *testing.T) {
	render := func(img string) string { return img + ".thumb" }

	cases := []struct {
		name    string
		images  []string
		workers int
		want    []string
	}{
		{"three_two_workers", []string{"a", "b", "c"}, 2, []string{"a.thumb", "b.thumb", "c.thumb"}},
		{"more_workers_than_images", []string{"z"}, 4, []string{"z.thumb"}},
		{"single_worker", []string{"c", "a", "b"}, 1, []string{"a.thumb", "b.thumb", "c.thumb"}},
		{"duplicate_images", []string{"x", "x"}, 3, []string{"x.thumb", "x.thumb"}},
		{"empty_batch", nil, 2, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RenderThumbnails(tc.images, tc.workers, render)
			if len(got) != len(tc.want) {
				t.Fatalf("RenderThumbnails(%v, %d) = %v, want %v", tc.images, tc.workers, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("RenderThumbnails(%v, %d) = %v, want %v", tc.images, tc.workers, got, tc.want)
				}
			}
		})
	}
}
