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

	for {
		var n int
		_, err := fmt.Fscan(rd, &n)
		if err != nil {
			break
		}

		for i, t := 1, 1; ; i++ {
			t %= n
			if t == 0 {
				fmt.Fprintf(wr, "%d\n", i)
				break
			}
			t = t*10 + 1
		}
	}
}
