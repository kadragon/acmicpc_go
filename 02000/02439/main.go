package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	arr []byte

	n int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	arr = make([]byte, 101)
}

func f(k int) {
	if k < 1 {
		return
	}
	f(k - 1)
	arr[n-k+1] = '*'
	fmt.Fprintf(wr, "%s\n", string(arr[1:n+1]))
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		arr[i] = ' '
	}

	f(n)
}
