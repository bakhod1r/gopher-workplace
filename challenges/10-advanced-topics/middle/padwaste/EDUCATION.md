# How Many Bytes Is The Padding

## Intuition

Padding is not stored anywhere you can query directly — it is the difference between what the fields need and what the struct occupies. Reflection gives you both numbers.

## Approach

1. Reject non-structs.
2. Sum `t.Field(i).Type.Size()` over every field.
3. Return `t.Size()` minus that sum.

## Solution

```go
import "reflect"

// Waste returns how many bytes of v's struct type are padding: its total
// size minus the sum of its fields' sizes.
//
// A non-struct wastes nothing.
//
// Examples:
//
// 	Waste(gappy{}) => 14 for a byte, an int64 and a byte
func Waste(v any) uintptr {
	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Struct {
		return 0
	}
	var used uintptr
	for i := 0; i < t.NumField(); i++ {
		used += t.Field(i).Type.Size()
	}
	return t.Size() - used
}
```

## Walkthrough

`gappy` needs 1 + 8 + 1 = 10 bytes of fields and occupies 24: seven bytes after `A` and seven after `C`. `packed` occupies 16, wasting 6.

## Pitfalls

- Skipping unexported fields, which undercounts the content and overstates the waste.
- Using `Sizeof` on the interface value, which measures the interface header, not the struct.
