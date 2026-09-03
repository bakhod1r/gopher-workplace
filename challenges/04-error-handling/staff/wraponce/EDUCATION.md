# Idempotent Annotation

## Intuition

Layered code often annotates the same failure twice with the same word. Making the wrapper idempotent keeps messages readable without any layer needing to know what the others did.

## Approach

1. Return nil for nil.
2. Return `err` when its message already starts with `"<op>: "`.
3. Otherwise wrap with `%w`.

## Solution

```go
if err == nil {
	return nil
}
prefix := op + ": "
if strings.HasPrefix(err.Error(), prefix) {
	return err
}
return fmt.Errorf("%s%w", prefix, err)
```

## Walkthrough

A different operation name does not match the prefix, so `"outer: save: a"` is produced — only exact repetition is suppressed.

## Pitfalls

- Comparing without the separator, so `"saved: …"` is mistaken for the prefix.
- Re-wrapping on the skip path, breaking identity.
- Checking `strings.Contains`, which suppresses legitimate annotations.
