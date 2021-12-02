package common

import (
	"testing"

)

func TestGcdRecursiveWithAllPositive(t *testing.T) {
	m,n := 544,119
	want := 17
	actual,err := Gcd_Recursive(m,n)
	if (actual != want || err != nil) {
		t.Fatalf(`Gcd_Recursive(544, 119) = %q, %v. The wanted value is %#q`, actual, err, want)
	}
}