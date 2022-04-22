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

	var s string
	fmt.Fscan(rd, &s)

	var ans int
	for _, v := range s {
		ans = (ans*10 + int(v-'0')) % 20000303
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
