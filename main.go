package main

import (
	// access to stdin
	"os"
	// formatted output
	"fmt"
	// bufferizide in/output
	"bufio"
	// in/out
	"io"
)

func uniq(input io.Reader, output io.Writer) error {
	// read input full string -> Scanner
	in := bufio.NewScanner(input)
	// create map to write data that already occured
	// it works only for small amount of data in other case memory can be exhausted
	// alreadySeen := make(map[string]bool)

	// uniq approach
	var prev string

	// loop string by string
	// when it will be nothing to read scan will return false and we fall out of the loop
	for in.Scan() {
		txt := in.Text()

		// if key found skip this string
		// if _, found := alreadySeen[txt]; found {
		// 	continue
		// }

		// fill string
		// alreadySeen[txt] = true

		// compare by bytes if string is equal to previous
		if txt == prev {
			continue
		}

		// compare by bytes if data not sorted (somehow next string is smaller than prev)

		if txt < prev {
			return fmt.Errorf("data unsorted")
		}

		// assign et the end of loop
		prev = txt


		fmt.Fprintln(output, txt)
	}

	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	
	if err != nil {
		panic("some panic")
	}
}
