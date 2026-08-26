# Multi-Error Validation

## Intuition

Returning all errors at once (instead of the first one) gives the user a
complete picture. A nil slice means "no errors".

## Approach

1. Init `var errs []string` (nil).
2. Check each rule, append error message if failed.
3. Return `errs`.

## Solution

```go
func (u User) Validate() []string {
	var errs []string
	if u.Name == "" {
		errs = append(errs, "name is required")
	}
	if !strings.Contains(u.Email, "@") {
		errs = append(errs, "invalid email")
	}
	if u.Age < 0 {
		errs = append(errs, "age must be non-negative")
	}
	return errs
}
```

## Walkthrough

For `User{"", "bad", -1}`:
- Name empty → append "name is required".
- No "@" → append "invalid email".
- -1 < 0 → append "age must be non-negative".
- Returns 3-element slice.

## Pitfalls

- Initializing `errs := []string{}` instead of `var errs []string` — returns
  empty slice instead of nil when all valid.
- Returning `error` instead of `[]string` — this pattern returns *all* errors.
