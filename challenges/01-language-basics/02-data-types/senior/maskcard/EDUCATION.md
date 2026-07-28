# Index windows and inverted conditions

## The idea

The last four characters are indices `i >= len-4`; everything before is
`i < len-4`. To hide all but the last four, mask the **front**:

```go
if i < len(r)-4 { out[i] = '*' } else { out[i] = r[i] }
```

## Why it matters

Masking PII (cards, SSNs, tokens) is security-sensitive. An inverted window leaks
exactly the data you meant to hide while looking superficially "masked" — a real
data-exposure bug.

## Watch out

- Operate on runes so multi-byte input isn't miscounted.
- Return short inputs unchanged, or you might reveal everything (define the
  policy deliberately).
- The boundary is `len-4`; off-by-one reveals or hides one extra digit.
