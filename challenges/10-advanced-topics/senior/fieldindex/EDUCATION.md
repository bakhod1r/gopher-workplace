# Look The Field Up Once, Not Once Per Row

## Intuition

Every row in the slice has the same type, so the field's offset is decided once by the type, not once per row. `FieldByName` hands you that position — keep it and index directly.

## Approach

1. Validate the slice and its struct element type.
2. `FieldByName` once on the element type; check exported and int.
3. Keep `sf.Index` and use `FieldByIndex` inside the loop.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrShape is returned when rows is not a slice of structs with an int
// field of the given name.
var ErrShape = errors.New("rows must be a slice of structs with that int field")

// SumColumn totals the named int field across a slice of structs.
//
// Resolving the field by name is a string search through the struct's field
// table; doing it per row makes the cost O(rows x fields).
//
// Examples:
//
// 	SumColumn([]rec{{N: 1}, {N: 2}}, "N") => 3, nil
func SumColumn(rows any, field string) (int64, error) {
	rv := reflect.ValueOf(rows)
	if rv.Kind() != reflect.Slice {
		return 0, ErrShape
	}
	et := rv.Type().Elem()
	if et.Kind() != reflect.Struct {
		return 0, ErrShape
	}
	sf, ok := et.FieldByName(field)
	if !ok || !sf.IsExported() || sf.Type.Kind() != reflect.Int {
		return 0, ErrShape
	}
	idx := sf.Index
	var total int64
	for i := 0; i < rv.Len(); i++ {
		total += rv.Index(i).FieldByIndex(idx).Int()
	}
	return total, nil
}
```

## Walkthrough

For 4096 rows, `FieldByName` runs once and the loop does 4096 direct index lookups. Calling it per row would run 4096 string searches over the field table instead.

## Pitfalls

- Calling `rv.Index(i).FieldByName(field)` inside the loop — correct and slow.
- Validating the field inside the loop, so a bad shape is reported 4096 times.
