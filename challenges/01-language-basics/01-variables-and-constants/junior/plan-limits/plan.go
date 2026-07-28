package plan

// Tier identifies a customer subscription tier.
type Tier int

// TODO(candidate): declare the tier constants Free, Pro, Enterprise as
// distinct, ascending values (0, 1, 2) using iota.
const (
	Free       Tier = 0
	Pro        Tier = 0
	Enterprise Tier = 0
)

// Limit returns the requests-per-minute for a tier:
// Free => 60, Pro => 600, Enterprise => 6000. Unknown tier => 60.
//
// TODO(candidate): implement this. Map each tier to a multiple of 60
// (Pro = 60*10, Enterprise = 60*100) and default to 60.
func Limit(t Tier) int {
	panic("not implemented")
}
