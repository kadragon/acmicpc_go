package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
	d  []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = make([]int, 3)
}

func f(a, b, c int) (int, int) {
	if d[a]+d[b] == d[c] {
		return a, 0
	} else if d[a]-d[b] == d[c] {
		return a, 1
	} else if d[a]*d[b] == d[c] {
		return a, 2
	} else if d[a]/d[b] == d[c] {
		return a, 3
	} else if d[a]%d[b] == d[c] {
		return a, 4
	}
	return -1, -1
}

func main() {
	defer wr.Flush()
	flag := "+-*/%"

	fmt.Fscan(rd, &d[0], &d[1], &d[2])
	a, b := f(0, 1, 2)
	if a > -1 {
		wr.WriteString(fmt.Sprintf("%d%c%d=%d", d[0], flag[b], d[1], d[2]))
		return
	}
	_, b = f(1, 2, 0)
	wr.WriteString(fmt.Sprintf("%d=%d%c%d", d[0], d[1], flag[b], d[2]))
}
