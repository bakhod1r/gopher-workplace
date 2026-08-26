# Chain of Responsibility

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Handle` for `H20` in [chainresp.go](chainresp.go):

1. If `req == 20`, return `"twenty"`.
2. Otherwise, delegate to `h.Next(req)`.

## Validate

```bash
make verify
```
