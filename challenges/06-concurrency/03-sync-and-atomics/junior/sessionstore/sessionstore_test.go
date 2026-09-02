package sessionstore

import (
	"strconv"
	"sync"
	"testing"
)

func TestSessionStoreSaveLookup(t *testing.T) {
	cases := []struct {
		name   string
		saved  map[string]int
		token  string
		want   int
		wantOK bool
	}{
		{"active_session", map[string]int{"tok1": 7}, "tok1", 7, true},
		{"unknown_token", map[string]int{"tok1": 7}, "tok2", 0, false},
		{"reissued_token", map[string]int{"tok1": 9}, "tok1", 9, true},
		{"empty_store", map[string]int{}, "tok1", 0, false},
		{"user_zero", map[string]int{"tok1": 0}, "tok1", 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSessionStore()
			for k, v := range tc.saved {
				s.Save(k, v)
			}
			got, ok := s.Lookup(tc.token)
			if got != tc.want || ok != tc.wantOK {
				t.Errorf("Lookup(%q) = %d, %v, want %d, %v", tc.token, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

func TestSessionStoreConcurrent(t *testing.T) {
	s := NewSessionStore()
	const handlers = 16
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				token := strconv.Itoa((i*50 + j) % 20)
				s.Save(token, j)
				s.Lookup(token)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 20; i++ {
		if _, ok := s.Lookup(strconv.Itoa(i)); !ok {
			t.Errorf("Lookup(%q) missing after concurrent logins", strconv.Itoa(i))
		}
	}
}
