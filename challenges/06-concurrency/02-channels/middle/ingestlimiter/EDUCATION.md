# A Buffered Channel as a Semaphore

## Intuition

A buffered channel of capacity N is a bag of N tokens. Taking a token is a
send (it succeeds while there is room), and putting it back is a receive. When
all N are out, the next send blocks — which is exactly the queueing behaviour
you want in front of a service with a concurrency budget.

`struct{}` is the right element type: it occupies no memory, so the channel
holds permissions rather than data.

The second half of the puzzle is collecting results without a lock. Each worker
sends its result on a channel; the map is built by one goroutine — the caller —
so there is nothing to race over.

## Approach

1. Clamp `maxInFlight` to `>= 1`.
2. Build `slots` (capacity `maxInFlight`) and `results` (capacity `len(docs)`).
3. `wg.Add(len(docs))`; per document, start a goroutine that acquires, indexes,
   releases, and sends its result.
4. A closer goroutine: `wg.Wait()` then `close(results)`.
5. Range `results` into the map and return it.

## Solution

```go
// IndexDocuments indexes every document concurrently but never lets more than
// maxInFlight index calls run at the same time, because the index host falls
// over above its concurrency budget. A buffered channel of empty structs is
// the semaphore: acquiring a slot is a send, releasing it is a receive.
//
// The result maps document id to the score index returned. maxInFlight below 1
// is clamped to 1.
//
// Examples:
//
//	IndexDocuments([a b c], 2, score) => {a:.., b:.., c:..} with at most 2 concurrent calls
//	IndexDocuments([a], 8, score)     => {a:..}
//	IndexDocuments(nil, 4, score)     => {}
func IndexDocuments(docs []Document, maxInFlight int, index func(Document) int) map[string]int {
	if maxInFlight < 1 {
		maxInFlight = 1
	}

	type scored struct {
		id    string
		score int
	}

	slots := make(chan struct{}, maxInFlight)
	results := make(chan scored, len(docs))

	var wg sync.WaitGroup
	wg.Add(len(docs))
	for _, doc := range docs {
		go func(doc Document) {
			defer wg.Done()

			slots <- struct{}{}
			defer func() { <-slots }()

			results <- scored{id: doc.ID, score: index(doc)}
		}(doc)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	scores := map[string]int{}
	for r := range results {
		scores[r.id] = r.score
	}
	return scores
}
```

## Walkthrough

- With three documents and two slots, all three goroutines start, but the third
  parks on `slots <- struct{}{}` until one of the first two releases. Peak
  concurrency inside `index` is therefore 2.
- `results` is buffered to `len(docs)` so a worker never blocks on the send —
  it can release its slot promptly even before the caller starts collecting.
- The closer goroutine owns `close(results)`; the workers must not, since there
  are many of them.
- An empty batch adds 0 to the WaitGroup, so the closer closes immediately and
  the range yields an empty map.
- `maxInFlight = 0` clamps to 1, giving fully serialised indexing rather than a
  deadlock on a zero-capacity semaphore.

## Pitfalls

- A zero-capacity `slots` channel is not "unlimited", it is "nobody may proceed":
  the send blocks until a receiver appears, and there is none. Clamp first.
- Acquiring the slot before `go` serialises the *spawning*, not the work.
- Releasing without `defer` leaks a token whenever `index` panics, and the
  budget shrinks permanently.
- Writing into a shared map from every worker is a data race that `-race` will
  catch; route results through a channel or guard the map with a mutex.
- `wg.Wait()` on the caller's goroutine before ranging `results` deadlocks once
  `len(docs)` exceeds the buffer.
