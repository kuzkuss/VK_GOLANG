package main

import (
	"flag"
	"bufio"
	"fmt"
	"os"
	"errors"
	"github.com/kuzkuss/VK_GOLANG/uniq"
)

func main() {
	var opts uniq.Options
	uniq.Init(&opts)
	flag.Parse()

	inputStream := os.Stdin
	outputStream := os.Stdout

	var err error

	if len(flag.Args()) > 2 {
		fmt.Println(errors.New("error: too many args"))
		uniq.Usage()
		return
	}

	if filename := flag.Arg(0); filename != "" {
		inputStream, err = os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer inputStream.Close()
	}

	if filename := flag.Arg(1); filename != "" {
		outputStream, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outputStream.Close()
	}

	scanner := bufio.NewScanner(inputStream)
	
	strings, err := readStrings(scanner)

	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(outputStream)

	res, err := uniq.Uniq(strings, opts)

	if err != nil {
		fmt.Println(err)
		uniq.Usage()
		return
	}

	err = writeStrings(writer, res)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func readStrings(scanner *bufio.Scanner) ([]string, error) {
	strs := make([]string, 0)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	err := scanner.Err()
	return strs, err
}

func writeStrings(writer *bufio.Writer, strs []string) (error) {
	var err error
	for _, val := range strs {
		_, err = writer.WriteString(val + "\n")
		if err != nil {
			break
		}
	}
	return err
}
