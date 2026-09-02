# Formatter

## Intuition

This is the strategy pattern in its smallest form: the policy is a value, chosen by the caller, and the call site stays a single line.

## Approach

1. `Plain.Format` returns `msg`.
2. `KeyValue.Format` returns `"msg=" + msg`.
3. `Render` returns `f.Format(msg)`.

## Solution

```go
func (p Plain) Format(msg string) string { return msg }

func (k KeyValue) Format(msg string) string { return "msg=" + msg }

func Render(f Formatter, msg string) string { return f.Format(msg) }
```

## Walkthrough

`Render(KeyValue{}, "boom")` stores `KeyValue{}` in the interface. The dispatch lands on `KeyValue.Format`, producing `"msg=boom"`.

## Pitfalls

- Branching on the concrete type inside `Render` — the interface exists to avoid that.
- Special-casing the empty message; `"msg="` is the correct output.
- Returning a formatted copy of the wrong field name (`message=`).
