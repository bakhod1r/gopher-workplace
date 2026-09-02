package servercontext

import "testing"

func TestServerContext(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"not_nil", func(t *testing.T) {
			if ServerContext() == nil {
				t.Fatal("ServerContext() = nil, want non-nil context")
			}
		}},
		{"no_error", func(t *testing.T) {
			if err := ServerContext().Err(); err != nil {
				t.Errorf("ServerContext().Err() = %v, want nil", err)
			}
		}},
		{"done_is_nil", func(t *testing.T) {
			if ServerContext().Done() != nil {
				t.Error("ServerContext().Done() != nil, want a nil channel that never closes")
			}
		}},
		{"no_deadline", func(t *testing.T) {
			if _, ok := ServerContext().Deadline(); ok {
				t.Error("ServerContext().Deadline() ok = true, want false")
			}
		}},
		{"no_value", func(t *testing.T) {
			if v := ServerContext().Value("request-id"); v != nil {
				t.Errorf("ServerContext().Value = %v, want nil", v)
			}
		}},
		{"same_root_every_call", func(t *testing.T) {
			if ServerContext() != ServerContext() {
				t.Error("ServerContext() returned two different contexts, want one shared root")
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
