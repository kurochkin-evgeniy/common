package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stream := os.Stdin

	if len(os.Args) >= 2 {

		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		stream = file
	}

	m := make(map[string]int)

	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		m[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for key, val := range m {
		if val > 1 {
			fmt.Println(val, " # ", key)
		}
	}

}
