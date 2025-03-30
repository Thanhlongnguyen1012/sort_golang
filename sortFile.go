package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

// funtions: split files and organize
func sortFile(inputFile string) {
	const (
		fileNumber = 20
		number     = 200_000_000
	)
	// Create slice
	var slice []uint64
	var files [fileNumber]*os.File
	var writers [fileNumber]*bufio.Writer
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
		files[i] = f
		//bufer = 16 MB
		writers[i] = bufio.NewWriterSize(f, 16*1024*1024)
	}
	//Read input file, write slice = 250_000_000, write output file
	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("Read file error")
		}
		//Read an input file, store the data in a slice, then write it to an output file
		slice = append(slice, num)
		// if len(slice) = 250_000_000 then sort, Reset len of slice
		if len(slice) >= number {
			slices.Sort(slice)
			for line, num := range slice {
				_, err = writers[index].WriteString(strconv.FormatUint(num, 10) + "\n")
				if err != nil {
					fmt.Printf("error write line %d file %d"+"\n", line, index)
				}
			}
			index++
			//len = 0
			slice = slice[:0]
		}
	}
	//close file
	file.Close()
	for i := 0; i < fileNumber; i++ {
		writers[i].Flush()
		files[i].Close()
	}
	fmt.Println("Sort done !")
}
