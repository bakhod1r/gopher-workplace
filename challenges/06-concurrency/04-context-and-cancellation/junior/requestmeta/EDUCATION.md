# Tenant and Trace Together

## Intuition

Each `WithValue` adds exactly one link to the chain, so carrying two values means two wrappers — and a lookup walks outward-in, which is why the most recent middleware wins. Two distinct empty struct types cost no memory and make the two slots impossible to confuse, even though both hold a `string`.

## Approach

1. `WithMeta`: wrap `ctx` with the tenant, then wrap the result with the trace, and return it.
2. `Meta`: assert each lookup to `string`, discarding `ok` so a miss yields `""`.
3. Return both.

## Solution

```go
// WithMeta returns a copy of ctx carrying both the tenant ID and the trace ID
// that the edge middleware extracted from the request.
//
// Examples:
//
//	Meta(WithMeta(bg, "acme", "4bf9"))  => "acme", "4bf9"
//	Meta(context.Background())          => "", ""
//	Meta(WithMeta(WithMeta(bg, "a", "1"), "b", "2")) => "b", "2"
func WithMeta(ctx context.Context, tenant, trace string) context.Context {
	ctx = context.WithValue(ctx, tenantKey{}, tenant)
	return context.WithValue(ctx, traceKey{}, trace)
}

// Meta reports the tenant and trace IDs carried by ctx. Missing values come
// back as empty strings.
func Meta(ctx context.Context) (tenant, trace string) {
	tenant, _ = ctx.Value(tenantKey{}).(string)
	trace, _ = ctx.Value(traceKey{}).(string)
	return tenant, trace
}
```

## Walkthrough

- `WithMeta(bg, "acme", "4bf9")` builds two nested wrappers over `bg`.
- `Meta` walks the chain once per key and finds each value.
- On a bare background context both lookups return `nil`, and the assertions yield the empty string.
- A context that only has the tenant key set returns that tenant and an empty trace — the keys are fully independent.

## Pitfalls

- Do not bundle both values into one struct under one key unless every reader wants both; independent keys stay independently overridable.
- Two different key *types* are required — two variables of the same type would collide.
- Never let a business rule depend on a context value that can silently be missing; scope-critical values like a tenant ID should also be validated where they are read.
