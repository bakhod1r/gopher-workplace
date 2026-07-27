// Package checkoutgrid — Gopher Workplace challenge.
package checkoutgrid

// SeatMap returns a fixed 7-row by 10-column seating grid with every valid
// {row,col} in taken marked true. Coordinates outside [0,7)×[0,10) are ignored.
//
// Examples:
//
//	SeatMap([][2]int{{0,0}})        => grid with [0][0]=true
//	SeatMap([][2]int{{1,2},{6,9}})  => grid with [1][2],[6][9]=true
//	SeatMap([][2]int{{7,0}})        => all false (row 7 out of range)
func SeatMap(taken [][2]int) [7][10]bool {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
