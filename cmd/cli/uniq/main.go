package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// params
	count := flag.Bool("c", false, "prefix lines by the number of occurences")
	sortOutput := flag.Bool("s", false, "sort output")
	files := flag.Args()
	flag.Parse()

	// model
	counts := make(map[string]int)

	// view
	output := []string{}

	// fill model
	if len(files) == 0 {
		mapLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
				continue
			}

			mapLines(f, counts)
			f.Close()
		}
	}

	// view model
	for line, n := range counts {
		occurence := ""

		if *count {
			occurence += fmt.Sprintf("%d", n)
		}

		output = append(output, fmt.Sprintf("%s\t%s", occurence, line))
	}

	if *sortOutput {
		parseStartInt := func(str string) int64 {
			num, err := strconv.ParseInt(
				strings.Split(str, "\t")[0],
				10,
				64,
			)

			if err != nil {
				fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
				return 0
			}

			return num
		}

		sort.Slice(output, func(i, j int) bool {
			fmt.Println(parseStartInt(output[j]))

			return parseStartInt(output[i]) < parseStartInt(output[j])
		})
	}

	for _, line := range output {
		fmt.Println(line)
	}
}

// mapLines load the input lines as keys and occurence as value
func mapLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
