# Validate With Rules

## Intuition

Validation is the canonical case for collecting rather than aborting: the user wants every problem in one round trip, and the rule name is what makes each one actionable.

## Approach

1. Range over the rules in order.
2. Wrap each failure with the rule name.
3. Join the collected failures.

## Solution

```go
var errs []error
for _, r := range rules {
	if err := r.Fn(v); err != nil {
		errs = append(errs, fmt.Errorf("%s: %w", r.Name, err))
	}
}
return errors.Join(errs...)
```

## Walkthrough

The middle rule still runs after the first fails, and the joined message lists the two failures in declaration order.

## Pitfalls

- Returning at the first failure, hiding later problems.
- Losing the rule name, leaving the message unactionable.
- Sorting or deduplicating, which breaks the declared order.
