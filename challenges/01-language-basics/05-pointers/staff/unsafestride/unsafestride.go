// Package unsafestride reads the i-th int32 of an array via unsafe pointer
// arithmetic. A planted bug advances by i BYTES instead of i elements, reading
// misaligned/wrong data.
package unsafestride

import "unsafe"

// At returns arr[i] using unsafe.Add for the offset.
func At(arr *[4]int32, i int) int32 {
	base := unsafe.Pointer(arr)
	// CHANGE CODE BELOW THIS LINE
	p := unsafe.Add(base, i)
	// CHANGE CODE ABOVE THIS LINE
	return *(*int32)(p)
}
