package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
	dt [36]int64
	n  int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dt = [36]int64{1, 1, 2, 5}
}

func f(k int) int64 {
	if dt[k] != 0 {
		return dt[k]
	}
	for i := 0; i < k/2; i++ {
		dt[k] += f(i) * f(k-(i+1)) * 2
	}
	if k%2 == 1 {
		dt[k] += f(k/2) * f(k/2)
	}
	return dt[k]
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)
	fmt.Fprintf(wr, "%d\n", f(n))
}
