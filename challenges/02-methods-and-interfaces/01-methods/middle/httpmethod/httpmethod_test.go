package httpmethod

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		name     string
		appName  string
		wantBody string
	}{
		{"myapp", "myapp", "OK: myapp"},
		{"test", "test-service", "OK: test-service"},
		{"empty", "", "OK: "},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := &App{Name: tc.appName}
			h := app.Handler()
			rec := httptest.NewRecorder()
			h(rec, httptest.NewRequest("GET", "/health", nil))
			if got := rec.Body.String(); got != tc.wantBody {
				t.Errorf("Handler response = %q, want %q", got, tc.wantBody)
			}
		})
	}
}
