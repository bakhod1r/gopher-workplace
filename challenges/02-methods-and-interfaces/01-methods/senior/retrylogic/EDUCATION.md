# Retry Methods

## Intuition

A wrapper method can encapsulate retry logic (loops, sleeps, backoff) so the
caller just gets a single success/failure result.

## Solution

```go
func (c *Client) DoWithRetry(maxAttempts int) error {
	var err error
	for i := 0; i < maxAttempts; i++ {
		err = c.Do()
		if err == nil {
			return nil
		}
	}
	return err
}
```
