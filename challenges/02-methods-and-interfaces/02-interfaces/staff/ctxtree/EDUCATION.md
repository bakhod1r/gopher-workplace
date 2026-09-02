# Cancellation Tree

## Intuition

This is how `context.WithCancel` works underneath: cancellation is push-based. The parent holds its children and walks them, so no node needs a goroutine watching its parent, and closing the channel wakes every waiter at once.

## Approach

1. `Cancel` takes the lock, returns early if already cancelled, sets the flag, closes `done`, and takes the children slice.
2. Release the lock *before* recursing so a child's `Cancel` never contends with the parent's lock.
3. Clearing `n.children` drops references so cancelled subtrees can be collected.
4. `Child` registers under the lock, or cancels the new node if the parent is already done.

## Solution

```go
func (n *Node) Child() *Node {
	c := &Node{done: make(chan struct{})}

	n.mu.Lock()
	alreadyCancelled := n.cancelled
	if !alreadyCancelled {
		n.children = append(n.children, c)
	}
	n.mu.Unlock()

	if alreadyCancelled {
		c.Cancel()
	}
	return c
}

func (n *Node) Cancel() {
	n.mu.Lock()
	if n.cancelled {
		n.mu.Unlock()
		return
	}
	n.cancelled = true
	close(n.done)
	children := n.children
	n.children = nil
	n.mu.Unlock()

	for _, c := range children {
		c.Cancel()
	}
}
```

## Walkthrough

Registering a child and cancelling it are separated by the lock boundary: the parent decides under its own lock whether the child joins the tree, and the born-cancelled path runs outside that lock so no nested locking is possible.

## Pitfalls

- `close(n.done)` without the `cancelled` guard — the second cancel panics.
- Recursing into children while holding the parent's lock, which invites deadlock as the tree grows.
- Never clearing `children`, so a long-lived root accumulates every scope ever created — the classic context leak.
