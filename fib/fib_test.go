package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	n, err := Fib(36)
	if err != nil {
		panic(err)
	}
	if n != 14930352 {
		panic(fmt.Errorf("N is %d", n))
	}
	n, err = Fib(2)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic(fmt.Errorf("N is %d", n))
	}
	_, err = Fib(-1)
	if err == nil {
		panic(fmt.Errorf("No err on neg"))
	}
}
