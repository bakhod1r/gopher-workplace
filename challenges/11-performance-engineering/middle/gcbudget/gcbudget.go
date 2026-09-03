// Package gcbudget — Gopher Workplace challenge.
package gcbudget

// CPUFraction returns the share of CPU time the collector consumed:
// gcNS across all cores over the wall-clock window times the core count. A
// non-positive window or core count gives 0.
//
// Examples:
//
//	CPUFraction(100_000_000, 1_000_000_000, 1) => 0.1
func CPUFraction(gcNS, windowNS int64, cores int) float64 {
	panic("not implemented")
}

// TuneGOGC suggests a new GOGC value to bring the GC CPU fraction under a
// target: GC cost is roughly inversely proportional to the heap headroom, so
// doubling GOGC roughly halves it. The suggestion is
// currentGOGC * (currentFraction / targetFraction), rounded to the nearest
// whole value, never below the current value and never above maxGOGC. A non-positive current GOGC or
// target reports false, and so does an already-satisfied budget — there is
// nothing to tune.
//
// Examples:
//
//	TuneGOGC(100, 0.20, 0.10, 1000) => 200, true
func TuneGOGC(currentGOGC int, currentFraction, targetFraction float64, maxGOGC int) (int, bool) {
	panic("not implemented")
}

// MemoryCost returns the live heap multiplier a GOGC value implies:
// 1 + GOGC/100, which is what the extra CPU headroom costs in RAM. A
// non-positive GOGC reports false.
//
// Examples:
//
//	MemoryCost(200) => 3, true
func MemoryCost(gogc int) (float64, bool) {
	panic("not implemented")
}
