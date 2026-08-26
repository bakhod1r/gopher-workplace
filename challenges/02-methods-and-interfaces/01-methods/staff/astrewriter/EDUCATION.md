# AST Rewriter

## Solution

```go
func (n *Node) Rewrite() {
	if n == nil { return }
	if n.Type == "Ident" && n.Val == "foo" {
		n.Val = "bar"
	}
	n.Left.Rewrite()
	n.Right.Rewrite()
}
```
