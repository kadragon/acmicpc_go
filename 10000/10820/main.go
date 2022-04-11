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

	for {
		t, _, err := rd.ReadLine()
		if err != nil {
			return
		}

		var a, b, c, d int

		for _, p := range t {
			if p == 32 {
				d++
			} else if 97 <= p && p <= 122 {
				a++
			} else if 65 <= p && p <= 90 {
				b++
			} else {
				c++
			}
		}

		fmt.Fprintf(wr, "%d %d %d %d\n", a, b, c, d)
	}
}
