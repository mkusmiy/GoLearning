package main

import "testing"
import "strings"

// Test runs a test suite against WordCount.
func Test(t *testing.T) {
	ok := true

	//s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	//r := rot13Reader{s}
	//io.Copy(os.Stdout, &r)

	for _, c := range testCases {
		got := rot13Reader{strings.NewReader(c.in)}
		//got2, _ := io.Copy(os.Stdout, got1)
		//got := fmt.Sprintf("%v", got2)
		if strings.NewReader(c.want) != got.r {
			ok = false
		}
		if !ok {
			t.Fatalf("FAIL\n rot13Reader(%v) = %#v\n want:  %#v",
				c.in, got, c.want)
			break
		}
		t.Logf("PASS\n rot13Reader(%v) = %#v\n", c.in, got)
	}
}

var testCases = []struct {
	in   string
	want string
}{
	{"Lbh penpxrq gur pbqr!", "You cracked the code!"},
	{"1", "1"},
}
