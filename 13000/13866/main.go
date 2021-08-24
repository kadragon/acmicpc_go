package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	data := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscan(rd, &data[i])
	}

	sort.Ints(data)
	a, b := data[3]+data[0], data[2]+data[1]
	if a > b {
		a = a - b
	} else {
		a = b - a
	}

	wr.WriteString(fmt.Sprintf("%d", a))
}
