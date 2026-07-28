package codes

import "testing"

func TestSequence(t *testing.T) {
	want := map[string]Class{
		"Info": 0, "Success": 1, "Redirect": 2, "ClientError": 3, "ServerError": 4,
	}
	got := map[string]Class{
		"Info": Info, "Success": Success, "Redirect": Redirect,
		"ClientError": ClientError, "ServerError": ServerError,
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("%s=%d; want %d", k, got[k], w)
		}
	}
}
