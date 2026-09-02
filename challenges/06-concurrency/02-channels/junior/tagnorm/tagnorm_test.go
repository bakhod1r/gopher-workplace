package tagnorm

import (
	"reflect"
	"testing"
)

func chanOf(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestNormalizeTags(t *testing.T) {
	cases := []struct {
		name string
		tags []string
		want []string
	}{
		{"pair", []string{"az", "eu"}, []string{"AZ", "EU"}},
		{"mixed_case", []string{"Prod"}, []string{"PROD"}},
		{"empty_stream", nil, []string{}},
		{"already_upper", []string{"US", "EU"}, []string{"US", "EU"}},
		{"digits", []string{"az1"}, []string{"AZ1"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := make(chan string, len(tc.tags)+1)
			NormalizeTags(chanOf(tc.tags...), out)

			got := []string{}
			for v := range out {
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NormalizeTags(%v) sent %#v, want %#v", tc.tags, got, tc.want)
			}
		})
	}
}
