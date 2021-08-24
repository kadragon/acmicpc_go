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
	fmt.Fscan(rd, &n)

	var check, test string
	fmt.Fscan(rd, &check)

	for i := 0; i < n; i++ {
		ans := 1
		c := 0
		v := 0
		fmt.Fscan(rd, &test)

		for j := 0; j < len(test) && check[c] != '*'; j++ {
			if test[j] == check[c] {
				c++
				v = j
			} else {
				ans = 0
				break
			}
		}

		c = len(check) - 1
		for j := len(test) - 1; j >= 0 && check[c] != '*' && ans != 0; j-- {
			if j <= v {
				ans = 0
				break
			}
			if test[j] == check[c] {
				c--
			} else {
				ans = 0
				break
			}
		}

		if ans == 1 {
			wr.WriteString("DA\n")
		} else {
			wr.WriteString("NE\n")
		}
	}
}
