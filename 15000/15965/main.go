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
	k := 7368788
	fmt.Fscan(rd, &n)

	d := make([]bool, k+1)

	for i := 2; i <= k; i++ {
		if !d[i] {
			n--
			if n == 0 {
				fmt.Fprintln(wr, i)
				return
			}

			d[i] = true
			for j := 2; i*j <= k; j++ {
				d[i*j] = true
			}
		}
	}
}
