# Case conversion

## Intuition

ASCII upper/lower differ by one bit (0x20). Use `unicode.ToUpper` / `ToLower` for
clarity, and track whether you are at the start of a word:

```go
start := true
for _, r := range s {
	if r == ' ' { start = true; ...append space; continue }
	if start { append ToUpper(r) } else { append ToLower(r) }
	start = false
}
```

## Approach

1. Track atStart, true at string start and after each space. 2. At a word start, uppercase a-z. 3. Elsewhere in a word, lowercase A-Z. 4. Spaces reset atStart.

## Solution

```go
func Title(s string) string {
	out := []byte(s)
	atStart := true
	for i := 0; i < len(out); i++ {
		c := out[i]
		if c == ' ' {
			atStart = true
			continue
		}
		if atStart {
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			}
			atStart = false
		} else {
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
		}
		out[i] = c
	}
	return string(out)
}
```

## Walkthrough

Title("GO is FUN"): 'G' start->G; 'O' mid->o; space; 'i'->I; 's'->s; space; 'F'->F; 'U'->u; 'N'->n -> "Go Is Fun".

## Pitfalls

- `strings.Title` is deprecated (Unicode-incorrect for many scripts); this ASCII
  exercise avoids that dependency.
- Build output with `strings.Builder` or `[]rune`, not by `+=` in a loop.
- Multiple/leading spaces change word boundaries — decide the spec.
