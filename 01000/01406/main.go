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

	var s string
	fmt.Fscanf(rd, "%s\n", &s)

	a, b := []byte(s), []byte{}

	var n int
	fmt.Fscanf(rd, "%d", &n)

	for i := 0; i < n; i++ {
		var t byte
		fmt.Fscanf(rd, "\n%c", &t)

		la, lb := len(a), len(b)

		if t == 'L' {
			if la > 0 {
				b = append(b, a[la-1])
				a = a[0 : la-1]
			}
		} else if t == 'D' {
			if lb > 0 {
				a = append(a, b[lb-1])
				b = b[0 : lb-1]
			}
		} else if t == 'B' {
			if la > 0 {
				a = a[0 : la-1]
			}
		} else if t == 'P' {
			var tp byte

			fmt.Fscanf(rd, " %c", &tp)
			a = append(a, tp)
		}
	}

	for i := 0; i < len(a); i++ {
		fmt.Fprintf(wr, "%c", a[i])
	}

	for i := len(b) - 1; i >= 0; i-- {
		fmt.Fprintf(wr, "%c", b[i])
	}
}
