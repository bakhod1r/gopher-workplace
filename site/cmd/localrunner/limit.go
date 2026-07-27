package main

// Concurrency cap for toolchain work.
//
// Each /run or /vet forks a Go build plus a test binary — with -race, several
// hundred MB and every core. Left unbounded, a handful of rapid clicks (or one
// stuck run) is enough to swamp the machine the learner is working on. So we
// admit a small number at a time and shed the rest immediately with 503 rather
// than queueing work nobody is waiting for any more.

import (
	"net/http"
	"runtime"
	"time"
)

// maxConcurrentRuns is deliberately small: this serves one person at a
// keyboard, not a fleet.
var maxConcurrentRuns = max(2, min(4, runtime.NumCPU()/2))

// admitWait is how long a request waits for a slot before being shed. Long
// enough to absorb a double-click, short enough that the UI never looks hung.
const admitWait = 2 * time.Second

type limiter struct{ slots chan struct{} }

func newLimiter(n int) *limiter { return &limiter{slots: make(chan struct{}, n)} }

// acquire takes a slot, waiting up to admitWait. Reports whether it succeeded.
func (l *limiter) acquire() bool {
	t := time.NewTimer(admitWait)
	defer t.Stop()
	select {
	case l.slots <- struct{}{}:
		return true
	case <-t.C:
		return false
	}
}

func (l *limiter) release() { <-l.slots }

// withLimit wraps a handler that shells out to the Go toolchain. Over capacity
// it answers 503 with a Report-shaped body, so the frontend renders it as a
// normal error instead of a fetch failure.
func (s *server) withLimit(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.limiter.acquire() {
			w.WriteHeader(http.StatusServiceUnavailable)
			writeJSON(w, report{Error: "runner busy: too many runs in flight — try again in a moment"})
			return
		}
		defer s.limiter.release()
		h(w, r)
	}
}
