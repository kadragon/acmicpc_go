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

	d := []int{1, 0, 0, 2}

	for _, v := range s {
		switch v {
		case 'A':
			d[0], d[1] = d[1], d[0]
		case 'B':
			d[0], d[2] = d[2], d[0]
		case 'C':
			d[0], d[3] = d[3], d[0]
		case 'D':
			d[1], d[2] = d[2], d[1]
		case 'E':
			d[1], d[3] = d[3], d[1]
		case 'F':
			d[2], d[3] = d[3], d[2]
		}
	}

	a, b := 0, 0
	for i := 0; i < 4; i++ {
		if d[i] == 1 {
			a = i + 1
		} else if d[i] == 2 {
			b = i + 1
		}
	}

	fmt.Fprintf(wr, "%d\n%d\n", a, b)
}
