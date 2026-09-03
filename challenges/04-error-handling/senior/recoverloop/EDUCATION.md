# Isolate Each Item

## Intuition

`defer` is function-scoped, not block-scoped. Isolating each iteration means giving each iteration its own function call, which is what bounds the damage to one item.

## Approach

1. For each item, call a closure that defers a recovery.
2. Have the closure return the item's error.
3. Collect the results and join them.

## Solution

```go
var errs []error
for _, item := range items {
	err := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v: %w", r, ErrPanic)
			}
		}()
		return h(item)
	}()
	if err != nil {
		errs = append(errs, err)
	}
}
return errors.Join(errs...)
```

## Walkthrough

The panic on item 2 is recovered inside that iteration's closure, so item 3 is still handled and the batch reports one failure.

## Pitfalls

- Writing `defer` in the loop body, which defers to function exit and aborts the batch.
- Recovering around the whole loop, so the first panic ends the run.
- Swallowing the recovered value without recording an error.
