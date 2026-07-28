# Clearing struct fields through a pointer

## The idea

A struct pointer reaches every field, including nested pointer fields, letting a helper reset an object in place.

## Why it matters

Resetting/recycling nodes (object pools, list clearing) mutates via pointers.

## Watch out

- Setting `n.Next = nil` drops the reference (helps GC).
- Value receiver would clear a copy only.
