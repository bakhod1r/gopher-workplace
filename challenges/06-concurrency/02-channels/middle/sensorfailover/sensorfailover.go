// Package sensorfailover — Gopher Workplace challenge.
package sensorfailover

// Reading is one temperature sample from one sensor on the line.
type Reading struct {
	Sensor  string
	Celsius float64
}

// MergeSensorStreams collects every reading from the primary and backup sensor
// gateways, returning once both streams are exhausted. A stream that closes
// early must stop competing in the select — otherwise its closed channel is
// ready forever and starves the stream that is still delivering.
//
// Interleaving between the two gateways is not defined; every reading appears
// exactly once. A nil gateway is treated as already finished.
//
// Examples:
//
//	MergeSensorStreams(chan a,b | chan c) => 3 readings
//	MergeSensorStreams(chan a | nil)      => 1 reading
//	MergeSensorStreams(closed | closed)   => no readings
func MergeSensorStreams(primary, backup <-chan Reading) []Reading {
	// TODO(candidate): implement this.
	panic("not implemented")
}
