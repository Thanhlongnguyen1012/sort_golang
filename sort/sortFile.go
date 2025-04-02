package sort

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

// split, sort input file wrirte file
func SortFile(inputFile string, fileNumber int, numberSplit int) {
	// Create slice
	var slice []uint64
	var files []*os.File
	var writers []*bufio.Writer
	//Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error open input file")
	}
	defer file.Close()
	// Create output file
	for i := 0; i < fileNumber; i++ {
		f, err := os.Create(fmt.Sprintf("number%d_sort.txt", i))
		if err != nil {
			fmt.Printf("Create file %d err"+"\n", i)
		}
		files = append(files, f)
		writers = append(writers, bufio.NewWriterSize(f, 16*1024*1024)) //bufer = 16 MB
	}
	//Read an input file, store the data in a slice, then write it to an output file
	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("Read file error")
		}
		slice = append(slice, num)
		// if len(slice) = number then sort, Reset len of slice
		if len(slice) >= numberSplit {
			slices.Sort(slice)
			for line, num := range slice {
				_, err = writers[index].WriteString(strconv.FormatUint(num, 10) + "\n")
				if err != nil {
					fmt.Printf("error write line %d file %d"+"\n", line, index)
				}
			}
			index++
			slice = slice[:0] //reset, len = 0
		}
	}
	file.Close() //close file
	for i := 0; i < fileNumber; i++ {
		writers[i].Flush()
		files[i].Close()
	}
	fmt.Println("Sort done !")
}
