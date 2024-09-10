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
func sum(filePath string, chunkSim map[int][]string) (int, error) {
	data, err := readFile(filePath)
	if err != nil {
		return 0, err
	}

	_sum := 0
	// for _, b := range data {
	// 	intb := int(b)
	// 	_sum += intb
	// 	chunkSim[intb] = append(chunkSim[intb], filePath)
	// }


	for i := 0; i < len(data); i += 100 {
		chunk := 0
		for j := 0; j < 100; j++ {
			intb := int(data[j])
			_sum += intb
			chunk += intb
		}
		chunkSim[chunk] = append(chunkSim[chunk], filePath)
	}
	return _sum, nil
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	chunkSim := make(map[int][]string)
	for _, path := range os.Args[1:] {
		_sum, err := sum(path, chunkSim)

		if err != nil {
			continue
		}

		totalSum += int64(_sum)

		sums[_sum] = append(sums[_sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

	// for chunk, files := range chunkSim {
	// 	if len(files) > 1 {
	// 		fmt.Printf("ALO %d: %v\n", chunk, files)
	// 	}
	// }
}
