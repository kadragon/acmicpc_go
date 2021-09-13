package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dt = make([]int, 1000001)
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	for i := 2; i <= n; i++ {
		dt[i] = 123456789
		if i%2 == 0 {
			if dt[i/2]+1 < dt[i] {
				dt[i] = dt[i/2] + 1
			}
		}
		if i%3 == 0 {
			if dt[i/3]+1 < dt[i] {
				dt[i] = dt[i/3] + 1
			}
		}
		if dt[i-1]+1 < dt[i] {
			dt[i] = dt[i-1] + 1
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n])

	for n > 0 {
		fmt.Fprintf(wr, "%d ", n)
		if n%3 == 0 && dt[n/3]+1 == dt[n] {
			n /= 3
		} else if n%2 == 0 && dt[n/2]+1 == dt[n] {
			n /= 2
		} else {
			n--
		}
	}
}
