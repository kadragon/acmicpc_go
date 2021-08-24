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

	var n, s int
	fmt.Fscan(rd, &n, &s)

	data := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &data[i])
		data[i] += data[i-1]
	}

	var i, j, ans int
	for i = 1; i <= n; i++ {
		for ; data[i]-data[j] >= s; j++ {
			if ans == 0 || ans > i-j {
				ans = i - j
			}
		}
	}

	wr.WriteString(fmt.Sprintf("%#v", ans))
}
