# Every Message In The Chain

## Intuition

Because each wrapper's message embeds the one below it, the chain is a series of progressively shorter strings. Splitting them into separate fields is what makes them individually searchable.

## Approach

1. Declare `var out []string`.
2. While `err != nil`, append `err.Error()` and unwrap.
3. Return `out`.

## Solution

```go
var out []string
for err != nil {
	out = append(out, err.Error())
	err = errors.Unwrap(err)
}
return out
```

## Walkthrough

The two-wrap case yields three strings; the last is the bare root message with no prefixes.

## Pitfalls

- Initialising with `[]string{}`, so the nil case returns an empty slice.
- Unwrapping before appending, dropping the outermost message.
- Splitting `err.Error()` on `": "` instead of walking the chain.
