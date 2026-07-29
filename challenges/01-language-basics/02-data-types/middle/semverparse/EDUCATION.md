# Field-by-field numeric parsing

## Intuition

Parse a structured numeric string in one pass: accumulate digits with
`n = n*10 + (c-'0')`, and on a separator flush the field and move on. Validate
strictly — count fields, forbid empty ones and stray characters.

## Approach

1. Scan bytes, accumulating the current integer. 2. On '.', store the finished part, require at least one digit and fewer than 3 parts so far. 3. Non-digit chars are invalid. 4. At end require exactly two dots seen (idx==2) and a trailing digit.

## Solution

```go
func Parse(s string) (major, minor, patch int, ok bool) {
	parts := [3]int{}
	idx := 0
	digits := 0
	cur := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '.' {
			if digits == 0 || idx >= 2 {
				return 0, 0, 0, false
			}
			parts[idx] = cur
			idx++
			cur = 0
			digits = 0
			continue
		}
		if c < '0' || c > '9' {
			return 0, 0, 0, false
		}
		cur = cur*10 + int(c-'0')
		digits++
	}
	if idx != 2 || digits == 0 {
		return 0, 0, 0, false
	}
	parts[2] = cur
	return parts[0], parts[1], parts[2], true
}
```

## Walkthrough

Parse("1.4.10"): cur=1 at first '.' -> parts[0]=1; cur=4 -> parts[1]=4; cur=10 at end -> parts[2]=10. (1,4,10,true).

## Pitfalls

- Reject an empty field (`1..2`) and a trailing dot.
- Enforce the exact field count; `1.2` and `1.2.3.4` are both invalid here.
- A leading `-` or letters must fail — only `'0'..'9'` and `.` are valid.
