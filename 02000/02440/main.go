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
	fmt.Fprintf(wr, "%s\n", data[:n])
	f(n - 1)
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	f(n)
}
