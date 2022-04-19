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

	var a, b int
	fmt.Fscan(rd, &a, &b)

	if b > a {
		c := 100
		if b-a > 30 {
			c = 500
		} else if b-a > 20 {
			c = 270
		}
		fmt.Fprintf(wr, "You are speeding and your fine is $%d.\n", c)
	} else {
		fmt.Fprintf(wr, "Congratulations, you are within the speed limit!\n")
	}
}
