package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	cnt int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func f(d string) string {
	if len(d) == 1 {
		return d
	}

	var tmp int
	for i := 0; i < len(d); i++ {
		tmp += int(d[i]) - 48
	}

	cnt++
	return f(strconv.Itoa(tmp))
}

func main() {
	defer wr.Flush()

	var n string
	fmt.Fscan(rd, &n)

	n = f(n)

	fmt.Fprintf(wr, "%d\n", cnt)

	switch n[0] {
	case '3', '6', '9':
		fmt.Fprintf(wr, "YES")
	default:
		fmt.Fprintf(wr, "NO")
	}
}
