package main

import (
	"fmt"
	"sort_golang/create"
	"sort_golang/sort"
	"time"
)

func main() {
	start := time.Now()
	var numberTotal = 4_000_000_000
	var maxValue uint64 = 18_000_000_000
	var fileNumber = 20
	var numberSplit = numberTotal / fileNumber
	//create file Create a file of 4 ratios
	create.CreateFile("input.txt", numberTotal, maxValue)
	//split files and organize
	sort.SortFile("input.txt", fileNumber, numberSplit)
	var b []string
	for i := 0; i < fileNumber; i++ {
		b = append(b, fmt.Sprintf("number%d_sort.txt", i))
	}
	sort.MergeFile(b, "output.txt") //merger file
	delta := time.Now().Sub(start)
	fmt.Println("time: ", delta)
}
