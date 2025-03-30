package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		fileNumber = 20
	)
	start := time.Now()
	//create file Create a file of 4 ratios
	creaFile("input.txt")
	//split files and organize
	sortFile("input.txt")
	var b []string
	for i := 0; i < fileNumber; i++ {
		b = append(b, fmt.Sprintf("number%d_sort.txt", i))
	}
	//merger file
	mergeFile(b, "output.txt")
	delta := time.Now().Sub(start)
	fmt.Println("time: ", delta)
}
