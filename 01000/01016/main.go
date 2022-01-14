package main

import (
	"bufio"
	"fmt"
	"math"
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

	var min, max, ans int64
	fmt.Fscan(rd, &min, &max)
	ans = max - min + 1

	d := make([]bool, max-min+1)
	maxSquare := int64(math.Sqrt(float64(max)))

	for i := int64(2); i <= maxSquare; i++ {
		s := min / (i * i)
		if min%(i*i) > 0 {
			s++
		}
		for j := s; j*i*i <= max; j++ {
			t := j * i * i
			if !d[t-min] {
				ans--
				d[t-min] = true
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
