package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var n, p, count int
	var a string
	fmt.Fscan(rd, &n, &a)

	name := make([]string, n)
	answer := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &name[i], &answer[i])
		if strings.Compare(name[i], a) == 0 {
			p = i
			break
		}
	}

	for i := 0; i < p; i++ {
		if strings.Compare(answer[i], answer[p]) == 0 {
			count++
		}
	}

	fmt.Fprintf(wr, "%d\n", count)
}
