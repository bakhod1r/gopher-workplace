package scaleroundbug

import (
	"reflect"
	"testing"
	"time"
)

func TestScaleAllGrowth(t *testing.T) {
	got := ScaleAll([]int{7}, 300)
	if !reflect.DeepEqual(got, []int{21}) {
		t.Errorf("ScaleAll = %v, want [21]", got)
	}
}

func TestScaleAllRounding(t *testing.T) {
	got := ScaleAll([]int{250, 100}, 15)
	if !reflect.DeepEqual(got, []int{38, 15}) {
		t.Errorf("ScaleAll = %v, want [38 15]", got)
	}
}

func TestScaleAllNegative(t *testing.T) {
	got := ScaleAll([]int{-7, -250}, 300)
	if !reflect.DeepEqual(got, []int{-21, -750}) {
		t.Errorf("ScaleAll = %v, want [-21 -750]", got)
	}
}

func TestScaleAllScale(t *testing.T) {
	const n = 2000000
	vals := make([]int64, n)
	for i := range vals {
		vals[i] = 7
	}
	start := time.Now()
	got := ScaleAll(vals, 300)
	if el := time.Since(start); el > 5*time.Second {
		t.Fatalf("ScaleAll over %d elements took %v, want under 5s", n, el)
	}
	var sum int64
	for _, v := range got {
		sum += v
	}
	if want := int64(n) * 21; sum != want {
		t.Errorf("sum = %d, want %d", sum, want)
	}
}
