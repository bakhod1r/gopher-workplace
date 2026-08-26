# Comparison Methods

## Intuition

Go has no operator overloading. A `Compare` method returning -1/0/+1 is the
standard pattern (see `strings.Compare`, `cmp.Compare`).

## Approach

1. If `Major` differs, return based on that.
2. Otherwise compare `Minor`.

## Solution

```go
func (v Version) Compare(other Version) int {
	switch {
	case v.Major < other.Major:
		return -1
	case v.Major > other.Major:
		return 1
	case v.Minor < other.Minor:
		return -1
	case v.Minor > other.Minor:
		return 1
	default:
		return 0
	}
}
```

## Walkthrough

`Version{1, 5}.Compare(Version{1, 3})`:
- Major: 1 == 1 → continue.
- Minor: 5 > 3 → return 1.

## Pitfalls

- Forgetting to check Minor when Major is equal.
- Returning booleans instead of -1/0/+1.
