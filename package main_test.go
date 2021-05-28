package main

import "testing"

func TestMultiply(t *testing.T) {
	/* cmd: go test hellow\GoCodes\may */

	//Checking test result if enter number 3
	//if Multiply(3) != 4 {
	/*FAIL
	exit status 1
	FAIL    hellow/GoCodes/may      0.089s  */

	if Multiply(3) != 4 {
		t.Error("Expect 2 x 2 = 4")
		/* PASS
		ok      hellow/GoCodes/may      0.088s  */
	}
}
