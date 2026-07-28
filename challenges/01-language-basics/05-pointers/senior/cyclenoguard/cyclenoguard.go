// Package cyclenoguard detects a cycle with slow/fast pointers. A planted bug
// omits the fast.Next nil check, so an acyclic list panics on nil dereference.
package cyclenoguard

type Node struct {
	Val  int
	Next *Node
}

// HasCycle reports whether the list has a cycle.
func HasCycle(head *Node) bool {
	slow, fast := head, head
	// CHANGE CODE BELOW THIS LINE
	for fast != nil {
		// CHANGE CODE ABOVE THIS LINE
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
