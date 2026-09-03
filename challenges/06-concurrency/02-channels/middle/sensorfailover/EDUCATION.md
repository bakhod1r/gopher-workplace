# Disabling a Select Arm With Nil

## Intuition

Two facts collide here. A **closed** channel is always ready to receive — it
hands out zero values as fast as you ask. A **nil** channel is never ready — a
receive on it blocks forever. In a `select`, "never ready" is precisely what
"this case is switched off" means.

So when a gateway closes, you do not want to keep selecting on it; you want to
turn that arm off. Overwriting the local variable with `nil` does that, and the
loop condition then doubles as the termination test: when both arms are off,
there is nothing left to wait for.

## Approach

1. Start with an empty, non-nil result slice.
2. Loop while `primary != nil || backup != nil`.
3. `select` on both receives with comma-ok.
4. On `!ok`, set that variable to `nil` and `continue`.
5. Otherwise append the reading.
6. Return the collected readings.

## Solution

```go
// MergeSensorStreams collects every reading from the primary and backup sensor
// gateways, returning once both streams are exhausted. A stream that closes
// early must stop competing in the select — otherwise its closed channel is
// ready forever and starves the stream that is still delivering.
//
// Interleaving between the two gateways is not defined; every reading appears
// exactly once. A nil gateway is treated as already finished.
//
// Examples:
//
//	MergeSensorStreams(chan a,b | chan c) => 3 readings
//	MergeSensorStreams(chan a | nil)      => 1 reading
//	MergeSensorStreams(closed | closed)   => no readings
func MergeSensorStreams(primary, backup <-chan Reading) []Reading {
	readings := []Reading{}

	for primary != nil || backup != nil {
		select {
		case r, ok := <-primary:
			if !ok {
				primary = nil
				continue
			}
			readings = append(readings, r)
		case r, ok := <-backup:
			if !ok {
				backup = nil
				continue
			}
			readings = append(readings, r)
		}
	}

	return readings
}
```

## Walkthrough

- Both gateways live: the runtime picks a ready arm pseudo-randomly, so the
  interleaving varies — but each reading is delivered once and the totals match.
- The backup closes first: the next time its arm is chosen, `ok` is false, the
  variable becomes `nil`, and from then on only the primary arm can ever fire.
- Both closed: two more iterations turn both arms off and the loop condition
  ends the function.
- `backup == nil` on entry: that arm is dead from the first iteration; the loop
  still runs because the primary is live.
- Both `nil` on entry: the loop body never executes and `[]Reading{}` comes back.

## Pitfalls

- Without the nil trick, the closed gateway wins roughly half of all selects
  forever — the loop never terminates and burns a core.
- `break` inside a `select` case breaks the *select*, not the loop; a `for`/
  `select` termination needs a flag, a `return`, or this nil-arm technique.
- Dropping the comma-ok appends an endless stream of zero-valued `Reading`s.
- Setting the arm to a *closed* channel instead of nil changes nothing — closed
  is the problem, not the solution.
