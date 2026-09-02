package configswap

import (
	"sync"
	"testing"
)

func TestStore(t *testing.T) {
	cases := []struct {
		name    string
		stores  []Config
		want    Config
		wantVer int
	}{
		{"never_loaded", nil, Config{}, 0},
		{"first_publish", []Config{{1, "us"}}, Config{1, "us"}, 1},
		{"reload", []Config{{1, "us"}, {2, "eu"}}, Config{2, "eu"}, 2},
		{"three_reloads", []Config{{1, "us"}, {2, "eu"}, {3, "ap"}}, Config{3, "ap"}, 3},
		{"empty_region", []Config{{5, ""}}, Config{5, ""}, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s Store
			for _, c := range tc.stores {
				s.Store(c)
			}
			if got := s.Load(); got != tc.want {
				t.Errorf("Load() = %+v, want %+v", got, tc.want)
			}
			if got := s.Version(); got != tc.wantVer {
				t.Errorf("Version() = %d, want %d", got, tc.wantVer)
			}
		})
	}
}

func TestStoreConcurrentReload(t *testing.T) {
	var s Store
	s.Store(Config{Version: 1, Region: "r1"})

	var wg sync.WaitGroup
	wg.Add(21)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				c := s.Load()
				if c.Region != "r"+string(rune('0'+c.Version)) {
					t.Errorf("torn config: %+v", c)
					return
				}
			}
		}()
	}
	go func() {
		defer wg.Done()
		for v := 1; v <= 9; v++ {
			s.Store(Config{Version: v, Region: "r" + string(rune('0'+v))})
		}
	}()
	wg.Wait()
}
