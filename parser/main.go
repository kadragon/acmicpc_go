package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var rootDir string = "/Users/kadragon/Dev/acmicpc_go/"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please give one argument")
		return
	}

	sample := parser(args[0])
	writeFile(args[0], sample)
	// copyFile(args[0])
}

func parser(pid string) string {
	url := fmt.Sprintf("https://www.acmicpc.net/problem/%s", pid)

	c := colly.NewCollector(colly.Async(true))

	var result string
	c.OnHTML("#sampleinput1 pre", func(h *colly.HTMLElement) {
		result = h.Text
	})

	c.Visit(url)
	c.Wait()

	return result
}

func writeFile(pid, example string) error {
	folder := targetDir(pid)

	os.MkdirAll(folder, 0755)

	f, err := os.Create(folder + "/input.txt")
	if err != nil {
		return err
	}

	f.WriteString(example)
	f.Sync()

	defer f.Close()

	return nil
}

func targetDir(pid string) string {
	if len(pid) == 4 {
		pid = "0" + pid
	}

	return rootDir + pid[:2] + "000/" + pid
}
