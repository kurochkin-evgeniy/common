package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	"io/ioutil"
)

func getLineNumberForFile(fileName string) uint64 {
	file, err := os.Open(fileName)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineSum uint64
	for scanner.Scan() {
		lineSum++
	}

	return lineSum
}

func isFileSourceCode(fileName string) bool {
	extention := []string{"*.cpp","*.c", "*.hpp","*.h", "*.m", "*.mm", "*.cc", "*.sh"}

	for _, value := range extention {
		if match, _ := filepath.Match(value, fileName); match {
			return true
		}
	}

	return false
}

func isExcluded(fileName string) bool {
	skip := []string{"obj", "ext", "lib", "exe", "third-party", "boost"}

	for _, value := range skip {
		if value == fileName {
			return true
		}
	}

	// Skip hidden folders
	if match, _ := filepath.Match(".*", fileName); match {
		return true
	}

	return false
}

func getLineNumberForDirectory(dirName string) uint64 {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return 0
	}

	var lineSum uint64
	for _, file := range files {
		if !isExcluded(file.Name()) {
			fullPath := filepath.Join(dirName, file.Name())
			if file.IsDir() {
				lineSum += getLineNumberForDirectory(fullPath)
			} else {
				if isFileSourceCode(file.Name()) {
					lineSum += getLineNumberForFile(fullPath)
				} else {
					fmt.Println("Exclude file from count: ", fullPath)
				}
			}
		}
	}
	return lineSum
}

func main() {
	if len(os.Args) >= 2 {
		path := os.Args[1]
		fmt.Println(getLineNumberForDirectory(path))
	}
}
