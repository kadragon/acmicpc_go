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

	var x, y, t int
	for i := 0; i < 3; i++ {
		fmt.Fscanf(rd, "%d:", &t)
		if t > 0 && t <= 12 {
			x++
		}
		if t >= 0 && t < 60 {
			y++
		}
	}

	fmt.Fprintf(wr, "%d\n", x*(y-1)*(y-2))
}
