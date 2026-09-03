package poolscratch

import "testing"

func TestEncode(t *testing.T) {
	if got := Encode([]int{1, 2, 3}); got != "1,2,3" {
		t.Errorf("Encode = %q, want \"1,2,3\"", got)
	}
	if got := Encode(nil); got != "" {
		t.Errorf("Encode = %q, want empty", got)
	}
	if got := Encode([]int{-7}); got != "-7" {
		t.Errorf("Encode = %q, want \"-7\"", got)
	}
}

func TestEncodeRepeatedCallsStayCorrect(t *testing.T) {
	for i := 0; i < 100; i++ {
		if got := Encode([]int{i, i + 1}); got != itoa(i)+","+itoa(i+1) {
			t.Fatalf("call %d: Encode = %q", i, got)
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func TestEncodeUsesThePool(t *testing.T) {
	Encode([]int{1})
	got := pool.Get()
	if got == nil {
		t.Fatal("the pool is empty: the scratch buffer was never returned")
	}
	if b, ok := got.([]byte); !ok || cap(b) == 0 {
		t.Errorf("pooled value = %T, want a []byte with capacity", got)
	}
}
