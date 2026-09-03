// Package schedlatency — Gopher Workplace challenge.
package schedlatency

// Event is one goroutine's scheduling record from a trace: when it became
// runnable and when it actually started running.
type Event struct {
	G        int
	Runnable int64
	Running  int64
}

// Delay returns how long one goroutine waited for a core. An event whose
// Running is before its Runnable is malformed and reports false.
//
// Examples:
//
//	Delay(Event{1, 100, 150}) => 50, true
func Delay(e Event) (int64, bool) {
	panic("not implemented")
}

// Delays returns the valid delays in input order.
//
// Examples:
//
//	Delays([{1,0,10},{2,0,5}]) => []int64{10, 5}
func Delays(events []Event) []int64 {
	panic("not implemented")
}

// Stats returns the mean and the worst scheduling delay across the valid
// events, and false when there are none. A long tail here means the program
// has more runnable goroutines than cores — the CPU is saturated, and adding
// goroutines will make it worse.
//
// Examples:
//
//	Stats([{1,0,10},{2,0,20}]) => 15, 20, true
func Stats(events []Event) (mean float64, worst int64, ok bool) {
	panic("not implemented")
}
