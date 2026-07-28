# Mutable map values via pointers

## The idea

Map struct VALUES aren't addressable (`m[k].F = x` is illegal), so storing `*Struct` lets you fetch and mutate the bucket in place.

## Why it matters

Grouped accumulators and indexes use pointer-valued maps to update in place.

## Watch out

- `m[k].Total += x` on a map of struct VALUES won't compile.
- Store pointers and lazily create them.
