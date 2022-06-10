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

	var n, ans int

	for {
		fmt.Fscan(rd, &n)
		if n == -1 {
			break
		}
		ans += n
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
