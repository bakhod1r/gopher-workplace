package normalizegen

import (
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	if got := Normalize([]float64{2, 4}); !reflect.DeepEqual(got, []float64{0.5, 1}) {
		t.Errorf("Normalize([]float64{2, 4}) = %v, want [0.5 1]", got)
	}
	if got := Normalize([]float64{-4, 2}); !reflect.DeepEqual(got, []float64{-1, 0.5}) {
		t.Errorf("Normalize([]float64{-4, 2}) = %v, want [-1 0.5]", got)
	}
	if got := Normalize([]float64{0, 0}); !reflect.DeepEqual(got, []float64{0, 0}) {
		t.Errorf("Normalize([]float64{0, 0}) = %v, want [0 0]", got)
	}
	if got := Normalize([]float64{}); !reflect.DeepEqual(got, []float64{}) {
		t.Errorf("Normalize([]float64{}) = %v, want []", got)
	}
	if got := Normalize([]float32{1, 2}); !reflect.DeepEqual(got, []float32{0.5, 1}) {
		t.Errorf("Normalize([]float32{1, 2}) = %v, want [0.5 1]", got)
	}
}
