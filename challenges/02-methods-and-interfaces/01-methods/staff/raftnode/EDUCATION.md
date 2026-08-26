# Raft Node

## Solution

```go
func (n *Node) Timeout() {
	if n.State == "Follower" || n.State == "Candidate" {
		n.State = "Candidate"
	}
}

func (n *Node) ReceiveVotes(won bool) {
	if n.State == "Candidate" && won {
		n.State = "Leader"
	}
}
```
