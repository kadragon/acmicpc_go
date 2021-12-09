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

	for {
		var s string
		d := [26]int{}
		s, _ = rd.ReadString('\n')

		if len(s) == 2 && s[0] == '#' {
			break
		}

		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				d[c-'A']++
			} else if c >= 'a' && c <= 'z' {
				d[c-'a']++
			}
		}

		ans := 0
		for i := 0; i < 26; i++ {
			if d[i] > 0 {
				ans++
			}
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
