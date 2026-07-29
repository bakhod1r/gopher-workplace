# Sums and modular reduction

## Intuition

A simple checksum accumulates byte (or digit) values and reduces the total with
`%` to a fixed range:

```go
sum := 0
for _, b := range data { sum += int(b) }
return sum % 256
```

## Approach

1. Declare a uint8 accumulator (zero value 0).
2. Add each byte to it; uint8 arithmetic wraps modulo 256 automatically.
3. Return the accumulator.

## Solution

```go
func Checksum(data []byte) uint8 {
	var sum uint8
	for _, b := range data {
		sum += b
	}
	return sum
}
```

## Walkthrough

Checksum([]byte{200,100}): 0+200=200, 200+100=300 which wraps in uint8 to 44.

## Pitfalls

- The modulus defines the checksum width; pick it deliberately.
- Byte order and starting value matter for cross-implementation agreement.
- A plain sum catches single-byte errors but not reordering (that needs a
  position-weighted scheme like Luhn or Fletcher).
