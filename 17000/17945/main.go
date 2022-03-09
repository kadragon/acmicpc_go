package main

import (
	"bufio"
	"fmt"
	"math"
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

	var b, c int
	fmt.Fscan(rd, &b, &c)
	b = 2 * b

	t := int(math.Sqrt(float64(b*b - 4*c)))

	if b*b == 4*c {
		fmt.Fprintf(wr, "%d\n", (-b+t)/2)
	} else {
		fmt.Fprintf(wr, "%d %d\n", (-b-t)/2, (-b+t)/2)
	}
}
