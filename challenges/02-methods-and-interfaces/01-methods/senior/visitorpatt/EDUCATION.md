# Functional Visitor

## Solution

```go
func (n *Node) Accept(visitor func(int)) {
	if n == nil { return }
	visitor(n.Val)
	n.Left.Accept(visitor)
	n.Right.Accept(visitor)
}
```
