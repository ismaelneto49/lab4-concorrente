package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, ch chan int) {
	data, err := readFile(filePath)
	if err != nil {}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}
	ch <- _sum
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	ch := make(chan int)
	var totalSum int64
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		go sum(path, ch)
		// if err != nil {
		// 	continue
		// }

		// totalSum += int64(_sum)

		// sums[_sum] = append(sums[_sum], path)
	}
	var tsum int64
	for i := 0; i < 10; i++ { 
  	tsum += int64(<- ch)
	}
	
	fmt.Println(totalSum, "total")
	fmt.Println(tsum, "tsum")

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}
