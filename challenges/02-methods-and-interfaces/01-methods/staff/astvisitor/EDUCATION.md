# AST Visitor

## Solution

```go
func (n *Node) Visit(count *int) {
	if n == nil { return }
	if n.Type == "Ident" { *count++ }
	n.Left.Visit(count)
	n.Right.Visit(count)
}
```
