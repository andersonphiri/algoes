package common

import (
	"fmt"
)

func Gcd_Recursive(m, n int) (int, error) {
	if m < 0  {
		return -1, fmt.Errorf("the value of m %d should be at least %d", m,0)
	}
	if n < 0  {
		return -1, fmt.Errorf("the value of n %d should be at least %d", n,0)
	}
	if n >= m {
		return Gcd_RecursiveUtil(n, m)
	}
	return Gcd_RecursiveUtil(m, n)
}

func Gcd_RecursiveUtil(m, n int) (int, error) {
	r := m % n
	if r == 0 {
		return n, nil
	}
	return Gcd_RecursiveUtil(n, r)
}