package spinlock

import (
	"sync"
	"testing"
)

func TestLockUnlock(t *testing.T) {
	var l SpinLock
	if l.Locked() {
		t.Error("a fresh lock should not be held")
	}

	l.Lock()
	if !l.Locked() {
		t.Error("Locked = false after Lock")
	}
	if l.TryLock() {
		t.Error("TryLock should fail while the lock is held")
	}

	l.Unlock()
	if l.Locked() {
		t.Error("Locked = true after Unlock")
	}
	if !l.TryLock() {
		t.Error("TryLock should succeed on a free lock")
	}
	l.Unlock()
}

func TestUnlockOfUnlockedPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("unlocking an unlocked lock should panic")
		}
	}()

	var l SpinLock
	l.Unlock()
}

func TestMutualExclusion(t *testing.T) {
	var l SpinLock
	counter := 0

	const n = 1000
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++ // deliberately unsynchronised: the lock must protect it
			l.Unlock()
		}()
	}
	wg.Wait()

	if counter != n {
		t.Errorf("counter = %d, want %d", counter, n)
	}
	if l.Locked() {
		t.Error("the lock was left held")
	}
}

func TestTryLockUnderContention(t *testing.T) {
	var l SpinLock
	l.Lock()

	var wg sync.WaitGroup
	var mu sync.Mutex
	successes := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if l.TryLock() {
				mu.Lock()
				successes++
				mu.Unlock()
				l.Unlock()
			}
		}()
	}
	wg.Wait()

	if successes != 0 {
		t.Errorf("%d TryLocks succeeded while the lock was held", successes)
	}
	l.Unlock()
}

func TestSharedStateStaysConsistent(t *testing.T) {
	var l SpinLock
	data := make([]int, 0, 500)

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			l.Lock()
			data = append(data, i)
			l.Unlock()
		}(i)
	}
	wg.Wait()

	if len(data) != 500 {
		t.Errorf("len = %d, want 500 (appends were lost)", len(data))
	}
}

func TestIsLocker(t *testing.T) {
	var lk Locker = &SpinLock{}
	lk.Lock()
	lk.Unlock()
}

func BenchmarkSpinLockUncontended(b *testing.B) {
	var l SpinLock
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.Lock()
		l.Unlock()
	}
}

func BenchmarkMutexUncontended(b *testing.B) {
	var m sync.Mutex
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}
