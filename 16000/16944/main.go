package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var n, k int
	var s string
	fmt.Fscan(rd, &n, &s)
	k = n

	var number, lower, upper, special int

	for _, v := range s {
		if v >= '0' && v <= '9' {
			number++
		} else if v >= 'a' && v <= 'z' {
			lower++
		} else if v >= 'A' && v <= 'Z' {
			upper++
		} else {
			special++
		}
	}

	if number == 0 {
		k++
	}
	if lower == 0 {
		k++
	}
	if upper == 0 {
		k++
	}
	if special == 0 {
		k++
	}

	if k < 6 {
		k = 6
	}

	if k > n {
		fmt.Fprintf(wr, "%d\n", k-n)
	} else {
		fmt.Fprintf(wr, "0\n")
	}
}
