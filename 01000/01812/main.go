package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n int
	d []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func f(numberOfCandies int) bool {
	startCandies := numberOfCandies
	for i := 0; i < n; i++ {
		if d[i]-numberOfCandies < 0 {
			return false
		}
		numberOfCandies = d[i] - numberOfCandies
	}
	return startCandies == numberOfCandies
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 0; i <= d[0]; i++ {
		if f(i) {
			startCandies := i
			for j := 0; j < n; j++ {
				fmt.Fprintf(wr, "%d\n", startCandies)
				startCandies = d[j] - startCandies
			}
			return
		}
	}
}
