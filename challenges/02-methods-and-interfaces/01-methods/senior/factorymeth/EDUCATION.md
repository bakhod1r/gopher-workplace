# Factory Method

## Solution

```go
func (f StoreFactory) Create(storeType string) Store {
	switch storeType {
	case "mem":
		return MemStore{}
	case "disk":
		return DiskStore{}
	default:
		return nil
	}
}
```
