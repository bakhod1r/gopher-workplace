# State Machine

## Intuition

The state pattern's real content is the *transition table*: which states may
follow which. Putting it in one method makes illegal transitions unrepresentable
from the outside — callers can only say "advance", not "jump to Published".

## Approach

1. Switch on the current state.
2. Give each state its explicit successor.
3. Let the terminal state fall through unchanged.

## Solution

```go
func (d *Document) Publish() {
	switch d.State {
	case Draft:
		d.State = Moderation
	case Moderation:
		d.State = Published
	}
}
```

## Walkthrough

`Draft` is `iota == 0`. The first call matches the first arm and writes
`Moderation`. The second matches the second arm and writes `Published`. The
third matches no arm, so the switch does nothing — which is exactly the
"stays Published" requirement.

## Pitfalls

- **`d.State++`.** Tempting because the enum is contiguous, but it produces
  `State == 3`, an unnamed value with no meaning, on the third call. It also
  breaks the moment someone reorders the constants.
- **`if/else if` without a final else.** Equivalent, but the switch reads as the
  table it is.
- **Value receiver.** No transition is ever observed by the caller.
- **Exported `State` field.** Anyone can assign `Published` directly; a real
  implementation would unexport it and add a getter.
