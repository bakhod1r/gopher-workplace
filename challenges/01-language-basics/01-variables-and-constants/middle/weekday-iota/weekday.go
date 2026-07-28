// Package weekday enumerates days with iota.
package weekday

// Day is a day of the week, Sunday first.
type Day int

// Sunday..Saturday, Sunday=0.
//
// TODO(candidate): define the seven days with a single iota run.
const (
	Sunday    Day = 0
	Monday    Day = 0
	Tuesday   Day = 0
	Wednesday Day = 0
	Thursday  Day = 0
	Friday    Day = 0
	Saturday  Day = 0
)

// IsWeekend reports whether d is Saturday or Sunday.
//
// TODO(candidate): implement.
func IsWeekend(d Day) bool {
	panic("not implemented")
}
