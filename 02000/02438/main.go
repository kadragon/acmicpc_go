package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	data string
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	data = "****************************************************************************************************"
}

func f(n int) {
	if n < 1 {
		return
	}
	f(n - 1)
	fmt.Fprintf(wr, "%s\n", data[:n])
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	f(n)
}
