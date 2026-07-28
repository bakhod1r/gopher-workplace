# Mutating a struct's slice field

## The idea

A pointer-receiver method appends to and reassigns the struct's slice field, keeping the growth visible to callers.

## Why it matters

Stacks, queues, and buffers are structs with pointer-receiver mutators.

## Watch out

- A value receiver would append to a copy's field and lose it.
- Reassign the field with the append result.
