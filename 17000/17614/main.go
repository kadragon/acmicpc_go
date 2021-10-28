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
	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		for t := i; t > 0; t /= 10 {
			if t%10 == 3 || t%10 == 6 || t%10 == 9 {
				ans++
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
