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

	var n, ans int
	fmt.Fscan(rd, &n)

	for n > 0 {
		n >>= 1
		ans++
	}

	fmt.Fprintf(wr, "%d\n", 1<<(ans-1))
}
