# Spend A Budget

## Intuition

A budget makes the cost of a request predictable. Enforcing it before each unit of work — rather than counting afterwards — is what prevents the overspend from happening at all.

## Approach

1. Range over the steps.
2. Return `ErrBudgetExceeded` when the remaining budget is zero.
3. Decrement, run, and return any step failure immediately.

## Solution

```go
for _, step := range steps {
	if budget <= 0 {
		return ErrBudgetExceeded
	}
	budget--
	if err := step(); err != nil {
		return err
	}
}
return nil
```

## Walkthrough

With a budget of 1 and two steps the first runs, the second finds no budget left, and nothing beyond the allowance executes.

## Pitfalls

- Running the step and then checking the budget, spending one too many.
- Reporting the budget failure when a step already failed.
- Treating an empty step list with a zero budget as an error.
