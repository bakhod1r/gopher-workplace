# Boolean Methods

## Intuition

Methods that answer yes/no questions (`Is...`, `Has...`, `Can...`) are common in
Go. They read naturally: `person.IsAdult()`.

## Approach

1. Compare `Age` with 18.
2. Return the comparison result directly.

## Solution

```go
func (p Person) IsAdult() bool {
	return p.Age >= 18
}
```

## Walkthrough

For `Person{"Carol", 18}`:
- `18 >= 18` is `true`.

## Pitfalls

- Writing `if p.Age >= 18 { return true } else { return false }` — correct but
  verbose. Return the boolean expression directly.
- Off-by-one: `> 18` excludes 18-year-olds. The spec says `>= 18`.
