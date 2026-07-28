package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"}, {7, "7"},
	}
	for _, c := range cases {
		if got := FizzBuzz(c.n); got != c.want {
			t.Errorf("FizzBuzz(%d)=%q want %q", c.n, got, c.want)
		}
	}
}
