# Alignment vs size

## The idea

`Alignof` reports the address boundary a type must sit on; `Sizeof` reports its width. They match for basic types but the concepts (and APIs) are distinct.

## Why it matters

Layout and padding calculations depend on alignment, not size.

## Watch out

- Use `Alignof` for alignment, `Sizeof` for width.
- Struct padding is driven by field alignment.
