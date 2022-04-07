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

	d := [6][6]int{}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	var t int
	for k := 1; k <= 25; k++ {
		fmt.Fscan(rd, &t)

		for i := 1; i <= 5; i++ {
			for j := 1; j <= 5; j++ {
				if d[i][j] == t {
					d[i][j] = 0
					var c int
					for i := 1; i <= 5; i++ {
						var sx, sy int
						for j := 1; j <= 5; j++ {
							sx += d[i][j]
							sy += d[j][i]
						}
						if sx == 0 {
							c++
						}
						if sy == 0 {
							c++
						}
					}
					var sx, sy int
					for i := 1; i <= 5; i++ {
						sx += d[i][6-i]
						sy += d[i][i]
					}
					if sx == 0 {
						c++
					}
					if sy == 0 {
						c++
					}
					if c >= 3 {
						fmt.Fprintf(wr, "%d\n", k)
						return
					}
				}
			}
		}
	}
}
