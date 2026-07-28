// Package deferbeforeacq models acquire/release logging. A planted bug schedules
// the release defer BEFORE checking whether the resource was acquired, so a
// release is logged even when nothing was opened.
package deferbeforeacq

// Use logs "open" then "close" when ok is true, and logs NOTHING when ok is
// false (the resource was never acquired, so it must not be released).
func Use(ok bool) (log []string) {
	// CHANGE CODE BELOW THIS LINE
	defer func() { log = append(log, "close") }()
	if !ok {
		return
	}
	log = append(log, "open")
	// CHANGE CODE ABOVE THIS LINE
	return
}
