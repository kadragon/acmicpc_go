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
	fmt.Fscan(rd, &s)

	d := []int{0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2}

	con, rank, star := 0, 25, 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'W' {
			con++
			if con >= 3 && rank >= 6 {
				star += 2
			} else {
				star++
			}

			if star > d[rank] {
				star -= d[rank]
				rank--
			}
		} else {
			con = 0
			if rank < 21 {
				if star > 0 {
					star--
				} else {
					if rank < 20 {
						rank++
						star = d[rank] - 1
					}
				}
			}
		}

		if rank == 0 {
			fmt.Fprintf(wr, "Legend\n")
			return
		}
	}

	fmt.Fprintf(wr, "%d\n", rank)
}
