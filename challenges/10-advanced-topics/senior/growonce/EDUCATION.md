# One Allocation For A Stream Of Unknown Length

## Intuition

`io.ReadAll` cannot know the size, so it doubles. When the caller does know, the whole doubling chain collapses into one `make` — and the loop only needs a fallback for when the hint was wrong.

## Approach

1. Allocate `make([]byte, 0, hint+1)`.
2. Loop: if the buffer is full, grow it by appending one byte and reslicing back.
3. Read into the spare capacity and extend the length by what was read.
4. Return on `io.EOF`; propagate other errors.

## Solution

```go
import "io"

// Collect reads r to EOF and returns its bytes.
//
// hint is the caller's estimate of the size. When it is accurate the whole
// read must cost a single allocation instead of a chain of doublings.
//
// Examples:
//
// 	Collect(strings.NewReader("abc"), 3) => []byte("abc"), nil
func Collect(r io.Reader, hint int) ([]byte, error) {
	if hint < 0 {
		hint = 0
	}
	buf := make([]byte, 0, hint+1)
	for {
		if len(buf) == cap(buf) {
			buf = append(buf, 0)[:len(buf)]
		}
		n, err := r.Read(buf[len(buf):cap(buf)])
		buf = buf[:len(buf)+n]
		if err == io.EOF {
			return buf, nil
		}
		if err != nil {
			return buf, err
		}
	}
}
```

## Walkthrough

An 8192-byte payload with an exact hint allocates once and reads until EOF. `io.ReadAll` on the same payload allocates at 512, 1024, 2048, 4096 and 8192 bytes, copying each time.

## Pitfalls

- Reading into `buf[len(buf):]` — that window has length 0 when len < cap, so `Read` returns immediately and the loop spins.
- Discarding the bytes returned alongside `io.EOF`.
