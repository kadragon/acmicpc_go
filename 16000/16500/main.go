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

	t   string
	s   []string
	n   int
	ans int
)

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscanf(rd, "%s\n%d\n", &t, &n)

	s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%s\n", &s[i])
	}

	dt := [220]int{}

	for i := 0; i < len(t); i++ {
		for j := 0; j < n; j++ {
			if i == 0 || dt[i-1] == 1 {
				if strings.HasPrefix(t[i:], s[j]) {
					dt[i+len(s[j])-1] = 1
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[len(t)-1])
}
