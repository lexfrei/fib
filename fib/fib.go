package fib

import "fmt"

// Fib rerunt fib(c)
func Fib(c int) (n int, err error) {
	if c < 0 {
		return 0, fmt.Errorf("can't find fib for negative number. Use negfib")
	}

	if c == 0 {
		return 0, nil
	}

	var nm1, nm2 int

	nm1 = 1
	nm2 = 1

	for i := 0; i < c-2; i++ {
		nm2, nm1 = nm1, nm1+nm2
	}

	return nm1, nil
}
