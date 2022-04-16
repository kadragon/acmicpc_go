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

	var d, h, w int
	fmt.Fscan(rd, &d, &h, &w)

	a, b := h*h, w*w
	a, b = a*d*d/(a+b), b*d*d/(a+b)

	fmt.Fprintf(wr, "%d %d\n", int(math.Sqrt(float64(a))), int(math.Sqrt(float64(b))))
}
