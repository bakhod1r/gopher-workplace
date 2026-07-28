// Package popnoguard pops the head of a list-backed queue. A planted bug omits
// the empty check, dereferencing a nil head.
package popnoguard

type Node struct {
	Val  int
	Next *Node
}

type Queue struct{ head *Node }

// Pop removes and returns the front value and true, or 0,false when empty.
func (q *Queue) Pop() (int, bool) {
	// CHANGE CODE BELOW THIS LINE
	v := q.head.Val
	q.head = q.head.Next
	return v, true
	// CHANGE CODE ABOVE THIS LINE
}
