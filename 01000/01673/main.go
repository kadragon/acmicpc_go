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

	var n, k, ans int
	for {

		c, _ := fmt.Fscanf(rd, "%d %d\n", &n, &k)
		if c == 0 {
			return
		}

		ans = n

		for n/k > 0 {
			ans += n / k
			n = n/k + n%k
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
