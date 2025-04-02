package create

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// create input file
func CreateFile(s string, numberTotal int, maxValue uint64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	file, err := os.Create(s)
	if err != nil {
		fmt.Println("Error creating file")
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < numberTotal; i++ {
		num := r.Uint64() % (maxValue + 1)
		_, err = writer.WriteString(strconv.FormatUint(num, 10) + "\n")
		if err != nil {
			fmt.Printf("Error writing file on line %d"+"\n", i)
		}
	}
	writer.Flush()
	fmt.Println("Create done!")
}
