package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var s string
	for {
		fmt.Fscan(rd, &s)
		if s == "0" {
			break
		}

		fmt.Fprintf(wr, "%s\n", conv(s))
	}
}

func conv(s string) string {
	if len(s) == 1 {
		return s
	}

	t := 0
	for _, v := range s {
		t += int(v - '0')
	}

	return conv(strconv.Itoa(t))
}
