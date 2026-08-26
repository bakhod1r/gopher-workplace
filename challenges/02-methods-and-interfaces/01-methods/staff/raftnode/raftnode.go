// Package raftnode — Gopher Workplace challenge.
package raftnode

// Node represents a Raft node state.
type Node struct {
	State string // "Follower", "Candidate", "Leader"
}

// Timeout triggers an election if Follower or Candidate.
func (n *Node) Timeout() {
	// TODO(candidate): If Follower or Candidate, become Candidate.
	panic("not implemented")
}

// ReceiveVotes checks if majority won.
func (n *Node) ReceiveVotes(won bool) {
	// TODO(candidate): If Candidate and won, become Leader.
	panic("not implemented")
}
