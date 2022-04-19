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

	var n int
	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			return
		}

		var max, max_i int
		d := make([]string, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &d[i])

			c := 0

			for j := 1; j < len(d[i]); j++ {
				if d[i][j-1] == d[i][j] {
					for _, v := range "aeiouy" {
						if d[i][j-1] == byte(v) {
							c++
							break
						}
					}
				}
			}

			if max < c {
				max = c
				max_i = i
			}
		}

		fmt.Fprintf(wr, "%s\n", d[max_i])
	}
}
