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

	var t int
	fmt.Fscanf(rd, "%d\n", &t)
	for t > 0 {
		solve()
		t--
	}
}

func solve() {
	s, _ := rd.ReadString('\n')

	d := map[string]bool{}

	for {
		a, _ := rd.ReadString('\n')
		p := strings.Split(strings.TrimRight(a, "\n"), " ")
		if p[1] == "does" {
			break
		}
		d[p[2]] = true
	}

	for _, v := range strings.Split(strings.TrimRight(s, "\n"), " ") {
		if !d[v] {
			fmt.Fprintf(wr, "%s ", v)
		}
	}

	fmt.Fprintln(wr)
}
