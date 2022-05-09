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

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	s := readString()
	t := readString()

	p := len(t)

	ans := 0
	for {
		q := strings.Index(s, t)
		if q != -1 {
			ans++
			s = s[q+p:]
		} else {
			break
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}

func readString() string {
	s, _ := rd.ReadString('\n')
	return strings.TrimRight(s, "\n")
}
