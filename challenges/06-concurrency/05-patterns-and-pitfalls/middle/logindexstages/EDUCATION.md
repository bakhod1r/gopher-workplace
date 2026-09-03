# A Cancellable Three-Stage Pipeline

## Intuition

A pipeline is a chain of goroutines where each one reads from the previous
channel and writes to its own. Two rules keep it sane:

1. **The sender closes.** A stage closes only the channel it writes to. That
   close is the end-of-stream signal, and it cascades: `close(raw)` ends the
   parse stage's `range`, whose `defer close(parsed)` ends the caller's range.
2. **Every blocking send is a select.** Close propagation only travels
   *forwards*. When the consumer at the far end walks away, the upstream stages
   are still sitting on `ch <- v` with nobody receiving — that is a goroutine
   leak, and in a long-lived service it is a leak per batch. The `done` channel
   is the backwards signal, and `select` is how a send listens to it.

The filter is the easy part: a stage that receives N values and sends fewer is
just a `continue`.

## Approach

1. Make `raw`; start a goroutine that sends each line — `select` on `done` —
   and `defer close(raw)`.
2. Make `parsed`; start a goroutine that ranges over `raw`, calls `parse`,
   `continue`s when the record is not kept, sends otherwise — again `select` on
   `done` — and `defer close(parsed)`.
3. In the caller, range over `parsed`, bail out early if `done` is closed, and
   append `index(rec)`.

## Solution

```go
// IndexLogs streams raw log lines through a three-stage pipeline — read,
// parse-and-filter, index — and returns the index IDs in input order.
//
// Every stage selects on done as well as on its channels, so closing done
// tears the whole pipeline down: no stage is left blocked on a send that
// nobody will ever receive.
//
// Examples:
//
//	IndexLogs(open done, ["err disk", "info ok"], parse, index)  => ["idx:disk"]
//	IndexLogs(open done, ["info ok"], parse, index)              => nil
//	IndexLogs(closed done, ["err disk"], parse, index)           => nil
func IndexLogs(done <-chan struct{}, lines []string, parse func(string) (Record, bool), index func(Record) string) []string {
	raw := make(chan string)
	go func() {
		defer close(raw)
		for _, line := range lines {
			select {
			case raw <- line:
			case <-done:
				return
			}
		}
	}()

	parsed := make(chan Record)
	go func() {
		defer close(parsed)
		for line := range raw {
			rec, keep := parse(line)
			if !keep {
				continue
			}
			select {
			case parsed <- rec:
			case <-done:
				return
			}
		}
	}()

	var ids []string
	for rec := range parsed {
		select {
		case <-done:
			return ids
		default:
		}
		ids = append(ids, index(rec))
	}
	return ids
}
```

## Walkthrough

- **`["err disk", "info ok", "err io"]`.** The reader sends three lines. The
  parse stage keeps `err disk` and `err io`, drops `info ok` with `continue`.
  The caller appends `idx:disk` then `idx:io` — a single-channel chain carries
  one value at a time, so input order survives without any sorting.
- **Everything filtered out.** `parsed` receives nothing, is closed, the
  caller's `range` ends immediately and `ids` is still nil.
- **`done` already closed.** The reader's `select` may pick either arm on the
  first line, but every subsequent one returns; the caller's `default`-guarded
  check stops appending. Whatever the scheduler chooses, all three stages
  return — which is what the leak test hammers 50 times over.

## Pitfalls

- Closing `parsed` from the reader goroutine, or closing `raw` from the parse
  stage: whoever writes to a channel closes it, nobody else.
- A bare `raw <- line` with no `select`. It works right up until someone
  shuts the service down mid-batch, then leaks a goroutine per stage.
- Draining `raw` with `for { v := <-raw }` instead of `for v := range raw`
  spins forever on the zero value once the channel is closed.
- Reading the "keep" flag as an error flag. `parse` returning false is a
  routine dropped line, not a failure.
