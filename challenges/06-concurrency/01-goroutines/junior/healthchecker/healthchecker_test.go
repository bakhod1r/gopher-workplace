package healthchecker

import (
	"reflect"
	"testing"
)

func TestCheckAll(t *testing.T) {
	probe := func(service string) int {
		switch service {
		case "api":
			return 200
		case "db":
			return 503
		case "cache":
			return 301
		}
		return 500
	}

	cases := []struct {
		name     string
		services []string
		probe    func(string) int
		want     []bool
	}{
		{"up_and_down", []string{"api", "db"}, probe, []bool{true, false}},
		{"single_up", []string{"api"}, probe, []bool{true}},
		{"empty", []string{}, probe, []bool{}},
		{"redirect_is_healthy", []string{"cache"}, probe, []bool{true}},
		{"unknown_service", []string{"api", "ghost", "db"}, probe, []bool{true, false, false}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CheckAll(tc.services, tc.probe); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CheckAll(%v) = %v, want %v", tc.services, got, tc.want)
			}
		})
	}
}
