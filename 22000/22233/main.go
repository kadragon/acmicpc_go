package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n, t, cnt int
	var tmp string
	fmt.Fscan(rd, &n, &t)

	data := map[string]bool{}

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &tmp)
		
		data[tmp] = true
		cnt++
	}

	for i := 0; i < t; i++ {
		fmt.Fscan(rd, &tmp)
		for _, j := range strings.Split(tmp, ",") {
			if data[j] {
				cnt--
				data[j] = false
			}
		}

		fmt.Fprintf(wr, "%d\n", cnt)
	}
}
