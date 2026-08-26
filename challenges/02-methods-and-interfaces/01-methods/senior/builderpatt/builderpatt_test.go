package builderpatt

import "testing"

func TestBuilder(t *testing.T) {
	req := NewBuilder().
		Method("GET").
		URL("/api").
		Auth("token123").
		Build()

	if req.Method != "GET" || req.URL != "/api" || req.Auth != "token123" {
		t.Errorf("Build failed: %+v", req)
	}
}
