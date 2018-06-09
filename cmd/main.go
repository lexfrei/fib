package main

import (
	"flag"
	"fmt"

	"github.com/lexfrei/fib/fib"
)

var n int

func init() {
	flag.IntVar(&n, "n", 0, "Fib num")
	flag.Parse()
}

func main() {
	c, err := fib.Fib(n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d is %d\n", n, c)
}
