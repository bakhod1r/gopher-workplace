package nullobject

import "testing"

func TestRecorder(t *testing.T) {
	r := &Recorder{}
	r.Report("a")
	r.Report("b")
	if len(r.Events) != 2 || r.Events[1] != "b" {
		t.Errorf("Events = %v", r.Events)
	}
}

func TestNopMetrics(t *testing.T) {
	NopMetrics{}.Report("anything")
}

func TestMetricsOr(t *testing.T) {
	if _, ok := MetricsOr(nil).(NopMetrics); !ok {
		t.Error("MetricsOr(nil) should return NopMetrics")
	}
	r := &Recorder{}
	if got := MetricsOr(r); got != Metrics(r) {
		t.Error("MetricsOr should pass a non-nil value through")
	}
}

func TestProcessWithRecorder(t *testing.T) {
	r := &Recorder{}
	if got := Process(r, []string{"a", "b"}); got != 2 {
		t.Errorf("Process = %d, want 2", got)
	}
	if len(r.Events) != 2 || r.Events[0] != "item:a" {
		t.Errorf("Events = %v, want [item:a item:b]", r.Events)
	}
}

func TestProcessWithNil(t *testing.T) {
	if got := Process(nil, []string{"a"}); got != 1 {
		t.Errorf("Process = %d, want 1", got)
	}
	if got := Process(nil, nil); got != 0 {
		t.Errorf("Process = %d, want 0", got)
	}
}
