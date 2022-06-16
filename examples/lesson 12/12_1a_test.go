// To run the test use: go test 12_1a.go 12_1a_test.go
package main

import "testing"

func TestChangeCase(t *testing.T) {
	{
		s, expect := "This is it!", "tHIS IS IT!"
		if actual := changeCase(s); actual != expect {
			t.Errorf("%q: actual %q expect %q.", s, actual, expect)
		}
	}
}

func TestChangeCase0(t *testing.T) {
	{
		s, expect := "This is it!", "tHIS IS IT!"
		if actual := changeCase0(s); actual != expect {
			t.Errorf("%q: actual %q expect %q.", s, actual, expect)
		}
	}
}

func TestChangeCase1(t *testing.T) {
	{
		s, expect := "This is it!", "tHIS IS IT!"
		if actual := changeCase1(s); actual != expect {
			t.Errorf("%q: actual %q expect %q. %t", s, actual, expect, actual != expect)
		}
	}
}

func TestChangeCase2(t *testing.T) {
	{
		s, expect := "This is it!", "tHIS IS IT!"
		if actual := changeCase2(s); actual != expect {
			t.Errorf("%q: actual %q expect %q.", s, actual, expect)
		}
	}
}
