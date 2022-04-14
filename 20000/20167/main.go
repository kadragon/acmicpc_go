package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	max, n, k int
	d         []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n, &k)

	d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}
	f(0, 0, 0)

	fmt.Fprintf(wr, "%d\n", max)
}

func f(depth, sum, prev int) {
	if depth == n {
		if sum > max {
			max = sum
		}
		return
	}

	prev += d[depth]
	if prev >= k {
		f(depth+1, sum+prev-k, 0)
	} else {
		f(depth+1, sum, prev)
	}

	f(depth+1, sum, 0)
}
