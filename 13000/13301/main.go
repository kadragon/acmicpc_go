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

	dt := [81]int64{0, 1, 1}
	for i := 2; i <= n; i++ {
		dt[i] = dt[i-1] + dt[i-2]
	}
	dt2 := [81]int64{0, 4}
	for i := 2; i <= n; i++ {
		dt2[i] = dt2[i-1] + dt[i]*2
	}

	fmt.Fprintf(wr, "%d\n", dt2[n])
}
