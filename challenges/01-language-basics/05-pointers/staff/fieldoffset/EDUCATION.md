# Locating fields with unsafe.Offsetof

## The idea

`Offsetof` returns a field's byte position within its struct; `Sizeof` returns a type's width — different quantities.

## Why it matters

Manual field access (serialisation, cgo) must use Offsetof for positions.

## Watch out

- `Sizeof(p.B)` is 4 (width), `Offsetof(p.B)` is 4 (position) — they coincide here by luck of layout but mean different things; with a wider first field they diverge.
- Always use Offsetof for a field's position.
