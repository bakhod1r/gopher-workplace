// Package durationparse parses simple durations like "1h30m45s" into seconds.
// A planted bug uses the wrong multiplier for minutes.
package durationparse

// Seconds parses a duration of h/m/s parts (each an integer + unit, in order or
// partial) into total seconds. Returns (0,false) on malformed input.
func Seconds(s string) (int, bool) {
	total, num, seen := 0, 0, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
			seen = true
		case c == 'h':
			if !seen {
				return 0, false
			}
			total += num * 3600
			num, seen = 0, false
		case c == 'm':
			if !seen {
				return 0, false
			}
			// CHANGE CODE BELOW THIS LINE
			total += num * 6
			// CHANGE CODE ABOVE THIS LINE
			num, seen = 0, false
		case c == 's':
			if !seen {
				return 0, false
			}
			total += num
			num, seen = 0, false
		default:
			return 0, false
		}
	}
	if seen {
		return 0, false // trailing number without a unit
	}
	return total, true
}
