package sort

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// Create a struct containing the value,
// index of an element in the heap
type Item struct {
	value uint64
	index int
}

// Create min heap type use package container/heap
type minHeap []Item

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Merge files using minHeap structure
// Write output file
func MergeFile(inputFile []string, outputFile string) {
	files := make([]*os.File, len(inputFile)) // open input file
	scanners := make([]*bufio.Scanner, len(inputFile))
	//open input file
	for i, fileName := range inputFile {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Opening the %d file failed ", i)
		}
		defer file.Close()
		files[i] = file
		scanners[i] = bufio.NewScanner(file)
	}
	fileOut, err := os.Create(outputFile) // open output file
	if err != nil {
		println("Opening the output file failed")
	}
	defer fileOut.Close()
	writer := bufio.NewWriter(fileOut)
	//initialize value for minHeap
	h := &minHeap{}
	heap.Init(h)
	for i, scanner := range scanners {
		if scanner.Scan() {
			num, err := strconv.ParseUint(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("Error converting to uint64")
			}
			heap.Push(h, Item{value: num, index: i})
		}
	}
	//Take the smallest element out of minHeap and add an element to minHeap
	//1 element goes into the heap, 1 element goes out of the heap
	for h.Len() > 0 {
		minItem := heap.Pop(h).(Item)
		writer.WriteString(strconv.FormatUint(minItem.value, 10) + "\n")
		if scanners[minItem.index].Scan() {
			num, err := strconv.ParseUint(scanners[minItem.index].Text(), 10, 64)
			if err != nil {
				fmt.Println("Error converting to uint64")
			}
			heap.Push(h, Item{value: num, index: minItem.index})
		}
	}
	writer.Flush()
}
