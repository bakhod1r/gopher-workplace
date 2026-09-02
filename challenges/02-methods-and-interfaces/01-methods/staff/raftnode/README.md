# Raft Node Transitions

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A Raft node is always in one of three roles. An election timeout pushes a
follower (or a candidate whose election stalled) into candidacy; winning a
majority promotes a candidate to leader. Encoding those transitions as methods
keeps illegal jumps — follower straight to leader — out of reach.

## Task

Implement `Timeout` and `ReceiveVotes` on `*Node` in [raftnode.go](raftnode.go):

1. `Timeout()`: a `"Follower"` or `"Candidate"` becomes `"Candidate"`. A
   `"Leader"` is unaffected.
2. `ReceiveVotes(won)`: a `"Candidate"` with `won == true` becomes `"Leader"`.
   Nothing else changes.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Follower; Timeout()
Output: Candidate
```

**Example 2:**

```
Input:  Candidate; ReceiveVotes(true)
Output: Leader
```

**Example 3:**

```
Input:  Follower; ReceiveVotes(true)
Output: Follower  (a non-candidate cannot win an election)
```

_Explanation:_ both the role *and* the event must match before a transition fires.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Guarded transitions** | Each method checks the current state before writing the next one. |
| 2 | **Leaders ignore timeouts** | A leader that re-elected itself on every tick would never make progress. |
| 3 | **Pointer receiver** | The role must persist between calls. |

## Hint

Two small guards. `Timeout` must **not** be an unconditional
`n.State = "Candidate"` — that would demote a healthy leader on every tick.

## Validate

```bash
make verify
```
