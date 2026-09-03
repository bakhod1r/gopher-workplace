# Verbose Formatting

## Intuition

Implementing `fmt.Formatter` moves the verbosity decision to the call site. The same value prints one line in production logs and a full trace under a debug flag.

## Approach

1. Return `Msg` from `Error`.
2. Switch on the verb inside `Format`.
3. Check `s.Flag('+')` for the verbose form and write to `s`.

## Solution

```go
switch verb {
case 'v':
	if s.Flag('+') {
		io.WriteString(s, e.Msg+"\n\t"+e.Detail)
		return
	}
	io.WriteString(s, e.Msg)
case 's':
	io.WriteString(s, e.Msg)
default:
	io.WriteString(s, e.Msg)
}
```

## Walkthrough

`fmt.Sprintf("%+v", e)` calls `Format` with verb `'v'` and the plus flag set, so the detail line is appended.

## Pitfalls

- Implementing only `Error`, so `%+v` prints the short message.
- Calling `fmt.Fprintf(s, "%v", e)` inside `Format`, which recurses forever.
- Ignoring the verb and printing the same thing for `%q` and `%d`.
