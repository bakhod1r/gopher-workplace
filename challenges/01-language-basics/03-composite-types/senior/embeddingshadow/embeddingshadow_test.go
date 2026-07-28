package embeddingshadow

import "testing"

func TestBaseID(t *testing.T) {
	e := Entity{Base: Base{ID: 42}, ID: 7}
	if got := BaseID(e); got != 42 {
		t.Errorf("BaseID=%d; want 42 (embedded Base.ID)", got)
	}
}
