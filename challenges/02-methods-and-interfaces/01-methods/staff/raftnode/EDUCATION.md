# Raft Node Transitions

## Intuition

Raft's safety argument rests on roles being hard to change: you become a
candidate only by timing out, and a leader only by winning votes. Writing the
transitions as guarded methods is how that argument survives contact with code —
no caller can assign `"Leader"` directly.

## Approach

1. `Timeout`: promote follower or candidate to candidate; leave a leader alone.
2. `ReceiveVotes`: promote a candidate to leader when it won; otherwise nothing.

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

## Walkthrough

The test starts a follower, times it out (→ candidate), then delivers a won
election (→ leader). A candidate timing out again stays a candidate — in real
Raft that is a new term with a fresh election, which is why the transition
targets candidate rather than being a no-op.

`ReceiveVotes(true)` on a follower does nothing: the role guard fails. That
guard is what keeps a stale vote response from promoting a node that has since
stepped down.

## Pitfalls

- **Unconditional assignment in `Timeout`.** Demotes a leader on every heartbeat
  tick — the cluster loses its leader continuously.
- **Ignoring `won`.** Any vote response, including a rejection, would make the
  node leader; that is a split-brain generator.
- **Checking `won` but not the role.** A follower or a leader would react to a
  vote reply meant for an election it is no longer running.
- **Value receiver.** No transition is ever observed.

## What real Raft adds

Terms. Every message carries a term number, and a node seeing a higher term
immediately steps down to follower — the rule that actually prevents two
leaders. This puzzle keeps the role machine alone so the guard pattern stays
visible.
