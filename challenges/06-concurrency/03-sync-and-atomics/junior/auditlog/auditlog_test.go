package auditlog

import (
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestAuditLog(t *testing.T) {
	cases := []struct {
		name    string
		appends []string
		want    string
		wantLen int
	}{
		{"empty", nil, "", 0},
		{"single", []string{"login"}, "login", 1},
		{"ordered", []string{"login", "logout"}, "login,logout", 2},
		{"duplicates", []string{"login", "login"}, "login,login", 2},
		{"three", []string{"a", "b", "c"}, "a,b,c", 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var l AuditLog
			for _, e := range tc.appends {
				l.Append(e)
			}
			got := l.Entries()
			if s := strings.Join(got, ","); s != tc.want {
				t.Errorf("Entries() = %q, want %q", s, tc.want)
			}
			if n := l.Len(); n != tc.wantLen {
				t.Errorf("Len() = %d, want %d", n, tc.wantLen)
			}
			if len(got) > 0 {
				got[0] = "tampered"
				if again := l.Entries(); again[0] == "tampered" {
					t.Error("Entries() returned the internal slice, not a copy")
				}
			}
		})
	}
}

func TestAuditLogConcurrent(t *testing.T) {
	var l AuditLog
	const handlers = 12
	const per = 100
	var wg sync.WaitGroup
	wg.Add(handlers + 2)
	for i := 0; i < handlers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < per; j++ {
				l.Append(strconv.Itoa(i) + "-" + strconv.Itoa(j))
			}
		}(i)
	}
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				for range l.Entries() {
				}
				l.Len()
			}
		}()
	}
	wg.Wait()

	if got, want := l.Len(), handlers*per; got != want {
		t.Errorf("Len() = %d, want %d", got, want)
	}
	if got, want := len(l.Entries()), handlers*per; got != want {
		t.Errorf("len(Entries()) = %d, want %d", got, want)
	}
}
