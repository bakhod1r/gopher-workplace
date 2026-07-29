# Call by value

## Intuition

Go copies each argument into the parameter; mutating a non-pointer parameter cannot affect the caller.

## Approach

1. Tax is `price*rate/100`.
2. Add it to the price.

## Solution

```go
func AddTax(price int, rate int) int {
	return price + price*rate/100
}
```

## Walkthrough

`AddTax(100, 20)`: 100 + 100*20/100 = 120.

## Pitfalls

- Multiply before dividing (`price*rate/100`) to avoid truncating rate/100 to 0.
- Slices/maps/pointers are the exceptions that DO share underlying data.
