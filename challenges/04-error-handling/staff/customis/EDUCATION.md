# Custom Is Method

## Intuition

`errors.Is` gives the error itself a say: if the error implements `Is`, that method decides the match. It is the hook for equivalence rules that plain value comparison cannot express.

## Approach

1. Assert the target to `*StatusError`, returning false otherwise.
2. Match an equal code directly.
3. Match a class marker by comparing `Code/100`.

## Solution

```go
t, ok := target.(*StatusError)
if !ok {
	return false
}
if t.Code == e.Code {
	return true
}
return t.Code%100 == 0 && t.Code/100 == e.Code/100
```

## Walkthrough

`errors.Is(&StatusError{503}, &StatusError{500})` calls the receiver's `Is`; 500 is a class marker and both codes divide to 5.

## Pitfalls

- Implementing `Is` on the target side, which is never consulted.
- Matching every `*StatusError`, so distinct codes collide.
- Treating 503 as a class marker, making the relation symmetric when it should not be.
