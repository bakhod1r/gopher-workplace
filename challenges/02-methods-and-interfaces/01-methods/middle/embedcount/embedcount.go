// Package embedcount — Gopher Workplace challenge.
package embedcount

// Tracker holds a count.
type Tracker struct {
	Count int
}

// Inc increments Count.
func (t *Tracker) Inc() {
	t.Count++
}

// Job embeds Tracker.
type Job struct {
	Tracker
	Name string
}

// Run simulates running the job by incrementing the embedded Tracker.
//
// Examples:
//
//	j := Job{Name: "task"}
//	j.Run() // j.Count == 1
func (j *Job) Run() {
	// TODO(candidate): call the promoted Inc() method.
	panic("not implemented")
}
