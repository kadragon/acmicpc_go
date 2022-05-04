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

	var n int
	fmt.Fscanf(rd, "%d\n", &n)

	for ; n > 0; n-- {
		s, _ := rd.ReadString('\n')

		t := strings.Split(s, " ")
		t[0] = "god"

		fmt.Fprintf(wr, "%s", strings.Join(t, ""))
	}
}
