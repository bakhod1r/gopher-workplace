package raftnode

import "testing"

func TestRaft(t *testing.T) {
	n := &Node{State: "Follower"}
	n.Timeout()
	if n.State != "Candidate" {
		t.Error("expected Candidate")
	}

	n.ReceiveVotes(true)
	if n.State != "Leader" {
		t.Error("expected Leader")
	}
}
