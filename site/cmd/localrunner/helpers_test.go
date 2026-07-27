package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRecorder() *httptest.ResponseRecorder { return httptest.NewRecorder() }

// captureLog redirects the standard logger into a buffer for the duration of a
// test, so log-only behaviour (warnings, best-effort errors) can be asserted.
func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	flags := log.Flags()
	out := log.Writer()
	log.SetOutput(&buf)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(out)
		log.SetFlags(flags)
	})
	return &buf
}

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	writeJSON(rec, map[string]string{"a": "b"})
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q", got)
	}
	if got := rec.Body.String(); got != "{\"a\":\"b\"}\n" {
		t.Errorf("body = %q", got)
	}
}

func TestDecodeBody(t *testing.T) {
	var v struct {
		Src string `json:"src"`
	}
	req := httptest.NewRequest(http.MethodPost, "/fmt", bytes.NewReader([]byte(`{"src":"x"}`)))
	if err := decodeBody(req, &v); err != nil || v.Src != "x" {
		t.Errorf("decodeBody = %v, %q", err, v.Src)
	}
	if err := decodeBody(httptest.NewRequest(http.MethodGet, "/fmt", nil), &v); err == nil {
		t.Error("GET accepted")
	}
	req = httptest.NewRequest(http.MethodPost, "/fmt", bytes.NewReader([]byte("{oops")))
	if err := decodeBody(req, &v); err == nil {
		t.Error("malformed body accepted")
	}
}
