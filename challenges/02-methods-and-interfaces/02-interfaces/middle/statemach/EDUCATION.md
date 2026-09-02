# State Machine

## Intuition

Modelling states as types puts each transition table next to the state it belongs to. Adding a state means adding a type, not editing a central switch.

## Approach

1. Each `Name` returns its label.
2. Each `Next` handles its own accepted events and otherwise returns the receiver with `false`.
3. `Run` reassigns `s` on every event and counts the accepted ones.

## Solution

```go
func (p Pending) Next(event string) (State, bool) {
	if event == "ship" {
		return Shipped{}, true
	}
	return p, false
}

func (s Shipped) Next(event string) (State, bool) {
	if event == "deliver" {
		return Delivered{}, true
	}
	return s, false
}

func (d Delivered) Next(event string) (State, bool) { return d, false }

func Run(s State, events []string) (string, int) {
	accepted := 0
	for _, e := range events {
		next, ok := s.Next(e)
		s = next
		if ok {
			accepted++
		}
	}
	return s.Name(), accepted
}
```

## Walkthrough

`["deliver", "ship"]` from `Pending`: `deliver` is not accepted (stays pending), then `ship` moves to `shipped` — final state `shipped`, one accepted event.

## Pitfalls

- Returning `nil` for a rejected event, which makes the next call panic.
- Counting every event instead of the accepted ones.
- Forgetting to assign the returned state back to `s`, freezing the machine.
