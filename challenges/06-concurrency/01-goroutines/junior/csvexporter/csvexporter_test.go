package csvexporter

import (
	"reflect"
	"testing"
)

func TestRenderRows(t *testing.T) {
	cases := []struct {
		name string
		rows [][]string
		sep  string
		want []string
	}{
		{"comma", [][]string{{"a", "b"}}, ",", []string{"a,b"}},
		{"empty_record", [][]string{{}}, ",", []string{""}},
		{"empty", [][]string{}, ",", []string{}},
		{"two_records", [][]string{{"x"}, {"y", "z"}}, ";", []string{"x", "y;z"}},
		{"tab_separated", [][]string{{"1", "2", "3"}}, "\t", []string{"1\t2\t3"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RenderRows(tc.rows, tc.sep); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RenderRows(%v) = %v, want %v", tc.rows, got, tc.want)
			}
		})
	}
}
