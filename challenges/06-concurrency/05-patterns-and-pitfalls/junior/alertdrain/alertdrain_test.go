package alertdrain

import "testing"

func TestCountAlerts(t *testing.T) {
	cases := []struct {
		name        string
		alerts      []string
		closeAlerts bool
		closeDone   bool
		want        int
	}{
		{"three_alerts", []string{"disk", "cpu", "mem"}, true, false, 3},
		{"single_alert", []string{"disk"}, true, false, 1},
		{"five_alerts", []string{"a", "b", "c", "d", "e"}, true, false, 5},
		{"no_alerts_stream_closed", nil, true, false, 0},
		{"shutdown_before_start", nil, false, true, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			alerts := make(chan string, len(tc.alerts))
			for _, a := range tc.alerts {
				alerts <- a
			}
			if tc.closeAlerts {
				close(alerts)
			}

			done := make(chan struct{})
			if tc.closeDone {
				close(done)
			}

			if got := CountAlerts(done, alerts); got != tc.want {
				t.Errorf("CountAlerts() = %d, want %d", got, tc.want)
			}
		})
	}
}
