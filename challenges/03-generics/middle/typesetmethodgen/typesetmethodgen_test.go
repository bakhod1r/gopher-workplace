package typesetmethodgen

import (
	"reflect"
	"testing"
)

func TestLabels(t *testing.T) {
	labels, total := Labels([]Status{200, 404})
	if !reflect.DeepEqual(labels, []string{"ok", "missing"}) {
		t.Errorf("labels = %v, want [ok missing]", labels)
	}
	if total != 604 {
		t.Errorf("total = %d, want 604", total)
	}
}

func TestLabelsSingle(t *testing.T) {
	labels, total := Labels([]Status{200})
	if len(labels) != 1 || labels[0] != "ok" || total != 200 {
		t.Errorf("Labels = %v, %d, want [ok], 200", labels, total)
	}
}

func TestLabelsEmpty(t *testing.T) {
	labels, total := Labels([]Status{})
	if labels == nil || len(labels) != 0 || total != 0 {
		t.Errorf("Labels(empty) = %v, %d, want [], 0", labels, total)
	}
}
