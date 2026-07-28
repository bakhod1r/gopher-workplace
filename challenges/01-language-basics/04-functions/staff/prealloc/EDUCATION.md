# make length vs capacity

## The idea

`make([]T, n)` fills n zero elements; `make([]T, 0, n)` reserves capacity but stays empty — mixing the two with append yields leading zeros.

## Why it matters

The length/capacity make-and-append confusion is a very common slice bug.

## Watch out

- `make([]T, n)` + append = n zeros then your data.
- Use `make([]T, 0, n)` to reserve, or index into a length-n slice.
