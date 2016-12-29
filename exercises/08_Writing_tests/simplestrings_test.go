package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// test that Tuesday is a weekday
func TestContains1(t *testing.T) {
	var day string = "Tuesday"
	if got := Contains(weekdays, day); got == false {
		t.Error("Expected true, but got ", got)
	}
}

// test that Sunday is not a weekday
func TestContains2(t *testing.T) {
	var day string = "Sunday"
	if got := Contains(weekdays, day); got == true {
		t.Error("Expected true, but got ", got)
	}
}

// test that an empty search string returns 0
func TestContains3(t *testing.T) {
	var day string = ""
	if got := Contains(weekdays, day); got == false {
		t.Error("Expected true, but got ", got)
	}
}

// test that the string Monday is not found in the empty string
func TestContains4(t *testing.T) {
	var day string = "Monday"
	var emptyString string
	if got := Contains(emptyString, day); got == true {
		t.Error("Expected false, but got ", got)
	}
}
