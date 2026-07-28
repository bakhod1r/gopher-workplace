// Package lenoffbyone counts nodes. A planted bug starts the counter at 1 even
// for an empty list, over-counting by one.
package lenoffbyone

type Node struct {
	Val  int
	Next *Node
}

// Length returns the number of nodes. nil is 0.
func Length(head *Node) int {
	// CHANGE CODE BELOW THIS LINE
	count := 1
	// CHANGE CODE ABOVE THIS LINE
	for n := head; n != nil; n = n.Next {
		count++
	}
	return count
}
