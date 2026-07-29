# The Luhn algorithm

## Intuition

Scanning right to left, leave the odd positions alone and double the even ones;
if a doubled digit exceeds 9, subtract 9 (same as summing its digits). The number
is valid when the grand total is divisible by 10.

## Approach

1. Empty -> false. 2. Walk right to left, doubling every second digit. 3. If a doubled digit exceeds 9, subtract 9. 4. Sum all digits. 5. Valid iff sum%10==0. Non-digit chars are invalid.

## Solution

```go
func Valid(s string) bool {
	if s == "" {
		return false
	}
	sum := 0
	double := false
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}
```

## Walkthrough

Valid("79927398713"): summing with alternating doubling gives 70; 70%10==0 -> true.

## Pitfalls

- Parity is by position **from the right**, so reverse or index carefully.
- "Subtract 9" is equivalent to summing the two digits of the doubled value.
- Reject non-digit characters up front; a stray space or letter is invalid.
