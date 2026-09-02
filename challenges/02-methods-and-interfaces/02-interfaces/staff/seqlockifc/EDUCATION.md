# Sequence Lock

## Intuition

A seqlock optimises for read-heavy data: readers take no lock at all and instead *verify* afterwards that no write overlapped. The odd/even counter encodes "a write is in flight" in a single atomic word.

## Approach

1. `Write` serialises writers with a mutex, bumps the sequence to odd, stores the fields, and bumps it back to even.
2. `Read` loads the sequence; an odd value means retry.
3. Read the fields, then load the sequence again — equal and even means no write overlapped.
4. Otherwise loop and try again.

## Solution

```go
func (s *SeqLock) Write(requests, errors int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seq.Add(1) // odd: write in progress
	s.reqs.Store(requests)
	s.errs.Store(errors)
	s.seq.Add(1) // even: stable
}

func (s *SeqLock) Read() Snapshot {
	for {
		before := s.seq.Load()
		if before%2 != 0 {
			continue
		}

		snap := Snapshot{
			Requests: s.reqs.Load(),
			Errors:   s.errs.Load(),
		}

		if s.seq.Load() == before {
			return snap
		}
	}
}
```

## Walkthrough

The test writes pairs satisfying `Requests == Errors*2`. A reader that captured `Requests` from one write and `Errors` from the next would break the invariant — the second sequence check is what makes that observable and retried.

## Pitfalls

- Skipping the second sequence load: the read is then optimistic with no validation at all.
- Checking only for evenness at the end, missing a write that started and finished between the field reads.
- Using plain (non-atomic) fields, which is a data race even though the algorithm looks correct on paper.
