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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	u := []int{0, 105, 100, 120, 100, 135}

	var n int
	dt := [21]int{}
	fmt.Fscan(rd, &dt[0], &n)

	for i := 0; i < n; i++ {
		for j := 1; j <= 5; j++ {
			if dt[i+j] < dt[i]*u[j]/100 {
				dt[i+j] = dt[i] * u[j] / 100
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n])
}
