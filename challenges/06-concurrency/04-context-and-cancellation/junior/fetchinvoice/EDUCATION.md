# Context as the First Parameter

## Intuition

Putting `ctx` first is not decoration: it makes cancellation impossible to forget at a call site, and it lets every layer stop as one when the request ends. Checking it before the domain rules also keeps the error meaningful — reporting "invoice not found" for a request nobody is listening to would send an on-call engineer hunting a data bug that does not exist.

## Approach

1. Guard on `ctx.Err()` and return it if non-nil.
2. Return `ErrNotFound` for `id <= 0`.
3. Otherwise format and return the invoice.

## Solution

```go
import "fmt"

// FetchInvoice loads one invoice from the billing store. Following the standard
// convention, the context is the first parameter and is checked before the
// lookup runs.
//
// It returns ErrNotFound for a non-positive ID.
//
// Examples:
//
//	FetchInvoice(live ctx, 7)       => "invoice-7", nil
//	FetchInvoice(live ctx, 0)       => "", ErrNotFound
//	FetchInvoice(cancelled ctx, 7)  => "", context.Canceled
func FetchInvoice(ctx context.Context, id int) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if id <= 0 {
		return "", ErrNotFound
	}
	return fmt.Sprintf("invoice-%d", id), nil
}
```

## Walkthrough

- With a live context and a positive ID, the lookup returns `"invoice-7"`.
- With a live context and `id == 0`, the domain sentinel comes back.
- With a cancelled context the function returns `context.Canceled` even when the ID is also invalid — the last test case pins that ordering.

## Pitfalls

- Never store a context in a struct field; pass it explicitly as the first argument.
- Never pass a nil context — use `context.TODO()` if you genuinely have none yet.
- Do not collapse `ErrNotFound` and `ctx.Err()` into one error; callers retry one and not the other.
