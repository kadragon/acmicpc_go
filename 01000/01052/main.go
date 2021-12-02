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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	tmp := n

	for {
		t := n
		place, cnt := -1, 0
		for i := 0; t > 0; i++ {
			if t%2 == 1 {
				cnt++
				if place == -1 {
					place = i
				}
			}

			t /= 2
		}

		if cnt <= m {
			break
		}

		n += (1 << place)
	}

	fmt.Fprintf(wr, "%d\n", n-tmp)
}
