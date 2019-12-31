package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sorter/algorithm"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var alg *string = flag.String("a", "quicksort", "Sort alg")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()

	buffer := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := buffer.ReadLine()
		// read the end of infile
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		// convert string to int
		str := string(line)
		value, err2 := strconv.Atoi(str)
		if err2 != nil {
			err = err2
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "alg =", *alg)
	}

	// step 1: read values
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read values: ", values)

	// step 2: sort
	t1 := time.Now()
	switch *alg {
	case "bubblesort":
		algorithm.BubbleSort(values)
	case "quicksort":
		algorithm.QuickSort(values)
	default:
		fmt.Println("Sorting algorithm", *alg, "is either unknown or unsupported.")
	}
	t2 := time.Now()
	fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

	// step 3: write values
	err = writeValues(values, *outfile)
	if err != nil {
		fmt.Println("Failed to close the output file ", *outfile)
	}
}
