package visitorifc

import "testing"

func tree() Node {
	return Section{
		Title: "outer",
		Children: []Node{
			Text{Content: "a b"},
			Section{
				Title:    "inner",
				Children: []Node{Text{Content: "c d e"}},
			},
		},
	}
}

func TestWordCounter(t *testing.T) {
	w := &WordCounter{}
	Walk(tree(), w)
	if w.Words != 5 {
		t.Errorf("Words = %d, want 5", w.Words)
	}

	w2 := &WordCounter{}
	Walk(Text{Content: "one"}, w2)
	if w2.Words != 1 {
		t.Errorf("Words = %d, want 1", w2.Words)
	}

	w3 := &WordCounter{}
	Walk(Text{Content: ""}, w3)
	if w3.Words != 0 {
		t.Errorf("Words = %d, want 0", w3.Words)
	}
}

func TestHeadingCollector(t *testing.T) {
	h := &HeadingCollector{}
	Walk(tree(), h)

	want := []string{"outer", "inner"}
	if len(h.Titles) != len(want) {
		t.Fatalf("Titles = %v, want %v", h.Titles, want)
	}
	for i := range want {
		if h.Titles[i] != want[i] {
			t.Errorf("Titles = %v, want %v", h.Titles, want)
		}
	}
}

func TestVisitOnlyText(t *testing.T) {
	h := &HeadingCollector{}
	Walk(Text{Content: "x"}, h)
	if len(h.Titles) != 0 {
		t.Errorf("Titles = %v, want empty", h.Titles)
	}
}
