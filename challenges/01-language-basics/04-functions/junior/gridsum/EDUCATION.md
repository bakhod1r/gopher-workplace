# Nested iteration over 2D data

## The idea

A slice of slices is ragged; nested `range` visits each element without assuming a fixed width.

## Why it matters

Matrices, grids, and adjacency lists all use this traversal.

## Watch out

- Don't assume all rows share a length; range each row independently.
- A nil or empty grid ranges zero times.
