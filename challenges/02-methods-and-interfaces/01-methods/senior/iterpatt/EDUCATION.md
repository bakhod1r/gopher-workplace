# Iterator on Linked Lists

## Solution

```go
func (it *ListIter) HasNext() bool {
	return it.current != nil
}

func (it *ListIter) Next() int {
	val := it.current.Val
	it.current = it.current.Next
	return val
}
```
