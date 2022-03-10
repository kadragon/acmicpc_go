package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var t int
	fmt.Fscan(rd, &t)

	for t > 0 {
		t--
		fmt.Fprintln(wr, solve())
	}
}

func solve() string {
	var n int
	fmt.Fscan(rd, &n)

	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Strings(d)

	for i := 0; i < len(d)-1; i++ {
		if strings.HasPrefix(d[i+1], d[i]) {
			return "NO"
		}
	}

	return "YES"
}
