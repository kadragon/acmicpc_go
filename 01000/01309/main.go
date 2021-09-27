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

	var n int
	fmt.Fscan(rd, &n)

	dt := make([]int, 100001)
	dt[1] = 3
	dt[2] = 7

	for i := 3; i <= n; i++ {
		dt[i] = (2*dt[i-1] + dt[i-2]) % 9901
	}

	fmt.Fprintf(wr, "%d", dt[n])
}
