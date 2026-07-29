# checkoutgrid — Seating Grid Marker

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

You are building a simple seat-hold service for a small theater. The seating arrangement is fixed: 7 rows (0–6) by 10 columns (0–9). When customers hold seats, you record them as a list of `{row, column}` coordinates. Your task is to mark those seats on a grid so customers can see which seats are taken.

## Task

Implement `func SeatMap(taken [][2]int) [7][10]bool` to return a seating grid where:
- Each `true` value marks a taken seat
- Each `false` value marks an empty seat
- Coordinates outside the valid range `[0,7)×[0,10)` are ignored
- Do not modify the function signature or tests

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  taken=[][2]int{{0,0}}
Output: grid[0][0]=true, rest false
```

**Example 2:**

```
Input:  taken=[][2]int{{1,2},{6,9}}
Output: grid[1][2]=true, grid[6][9]=true
```

**Example 3:**

```
Input:  taken=[][2]int{{7,0},{0,10},{-1,0}}
Output: all false
```

_Explanation:_ Every coordinate is out of the 7x10 bounds, so all are ignored.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|


## Hint

1. Declare a variable `var g [7][10]bool` to create your seating grid. It will automatically be all-false.
2. Use a `for` loop to iterate over the `taken` slice.
3. Inside the loop, guard against out-of-range coordinates before indexing.
4. Return the completed grid.

## Validate

```bash
make verify
```
