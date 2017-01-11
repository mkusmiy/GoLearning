package main

import "testing"

// Test runs a test suite against WordCount.
func Test(t *testing.T) {
	ok := true
	for _, c := range testCases {
		got := NSqrt(c.in)
		if c.want != got {
			ok = false
		}
		if !ok {
			t.Fatalf("FAIL\n NSqrt(%v) = %#v\n want:  %#v",
				c.in, got, c.want)
			break
		}
		t.Logf("PASS\n NSqrt(%v) = %#v\n", c.in, got)
	}
}

var testCases = []struct {
	in   float64
	want float64
}{
	{4, 2},
	{9, 3},
}
