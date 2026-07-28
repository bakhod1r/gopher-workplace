# Padding side and width unit

## The idea

Left-padding **prepends** fill so the original text sits at the right edge:

```go
return pad + s
```

The pad count is `width - len([]rune(s))`, measured in runes so alignment is
visual, not byte-based.

## Why it matters

Fixed-width reports, aligned tables, and zero-padded IDs depend on the correct
side and the correct width unit. Right-padding or byte-counting silently
misaligns non-ASCII or numeric columns.

## Watch out

- Measure width in runes; byte width misaligns multi-byte text.
- Return unchanged when already at/over width — never truncate here.
- `strings.Repeat` builds the fill run efficiently.
