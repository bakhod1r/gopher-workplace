# Proxy

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Do` in [proxyobj.go](proxyobj.go):

1. Check `p.role == "admin"`. If so, return `p.w.Do()`.
2. Otherwise return `"denied"`.

## Validate

```bash
make verify
```
