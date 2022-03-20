package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var n, k int
	fmt.Fscan(rd, &n, &k)

	var s string
	fmt.Fscan(rd, &s)

	p := strings.Split(s, ",")
	d := make([]int, 0)
	for _, v := range p {
		q, _ := strconv.Atoi(v)
		d = append(d, q)
	}

	for i := 1; i <= k; i++ {
		t := make([]int, n-i)
		for j := 0; j < len(d)-1; j++ {
			t[j] = d[j+1] - d[j]
		}
		d = t
	}

	p = make([]string, 0)
	for _, v := range d {
		p = append(p, strconv.Itoa(v))
	}

	fmt.Fprintf(wr, "%s\n", strings.Join(p, ","))
}
