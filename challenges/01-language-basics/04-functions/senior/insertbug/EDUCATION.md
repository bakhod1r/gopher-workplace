# Shifting to insert into a slice

## The idea

Inserting at i needs the suffix `xs[i:]` moved into `xs[i+1:]` (rightward); copying the other way shifts left and overwrites.

## Why it matters

Insert/delete shift-direction bugs corrupt ordered buffers and gap arrays.

## Watch out

- Insert shifts RIGHT: `copy(xs[i+1:], xs[i:])`.
- Delete shifts LEFT: `copy(xs[i:], xs[i+1:])`.
