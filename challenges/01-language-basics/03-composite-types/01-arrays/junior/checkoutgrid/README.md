# checkoutgrid — Seating Grid Marker

## Context

You are building a simple seat-hold service for a small theater. The seating arrangement is fixed: 7 rows (0–6) by 10 columns (0–9). When customers hold seats, you record them as a list of `{row, column}` coordinates. Your task is to mark those seats on a grid so customers can see which seats are taken.

## Task

Implement `func SeatMap(taken [][2]int) [7][10]bool` to return a seating grid where:
- Each `true` value marks a taken seat
- Each `false` value marks an empty seat
- Coordinates outside the valid range `[0,7)×[0,10)` are ignored
- Do not modify the function signature or tests

## Examples

**Example 1: Single taken seat**
```go
SeatMap([][2]int{{0, 0}})
// Returns a 7×10 grid with only [0][0] = true
```

**Example 2: Multiple taken seats**
```go
SeatMap([][2]int{{1, 2}, {6, 9}})
// Returns a 7×10 grid with [1][2] = true and [6][9] = true
```

**Example 3: Out-of-range coordinates are ignored**
```go
SeatMap([][2]int{{7, 0}, {0, 10}, {-1, 0}})
// Row 7 is out of range (rows are 0–6), column 10 is out of range (columns are 0–9),
// and row -1 is negative. All are ignored; returns all-false grid.
```

**Example 4: Empty list**
```go
SeatMap(nil)
// Returns a 7×10 grid with all seats false
```

## Topics to Master

- **Array value types:** In Go, arrays are fixed-size values; `[7][10]bool` is a 7×10 grid
- **Fixed-size arrays vs. slices:** Slices are dynamic (`[][2]int`); arrays are fixed (`[7][10]bool`)
- **Zero value of arrays:** An uninitialized array is automatically filled with the zero value (`false` for `bool`)
- **Bounds checking:** Always verify row and column are within valid ranges before indexing
- **Indexing:** Use `array[row][col]` syntax to access and assign values

## Hint

1. Declare a variable `var g [7][10]bool` to create your seating grid. It will automatically be all-false.
2. Use a `for` loop to iterate over the `taken` slice.
3. Inside the loop, guard against out-of-range coordinates before indexing.
4. Return the completed grid.

## Validate

Run the following command from this directory to verify your solution:

```bash
make verify
```

This will check formatting, run static analysis, and execute all test cases.
