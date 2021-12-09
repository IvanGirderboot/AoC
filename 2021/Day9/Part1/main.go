package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var height_map [10000]int64
	//var low_points [10000]bool
	var risk_level_total int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line_number := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "")

		for i, v := range line {
			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Print(err)
			}
			index := (line_number * 100) + i
			height_map[index] = num
		}
		line_number++

	}

	fmt.Printf("A total of %d displays were found.\n", len(height_map))
	fmt.Printf("%v\n", height_map)

	for i := range height_map {
		if determineLowPoint(&height_map, int64(i)) {
			risk_level_total += (height_map[i] + 1)
			fmt.Printf("Risk level Sum is now %d\n", risk_level_total)
		}
	}

}

func determineLowPoint(hm *[10000]int64, i int64) bool {

	up_index := i - 100
	down_index := i + 100
	left_index := i - 1
	right_index := i + 1

	//lowest_point := p

	// Check above (unless it's out of bounds)
	if up_index > 0 {
		if (*hm)[i] >= (*hm)[up_index] {
			return false
		}
	}

	// Check Below (unless it's out of bounds)
	if down_index < int64(len(*hm)-1) {
		if (*hm)[i] >= (*hm)[down_index] {
			return false
		}
	}

	modulo := i % 100

	// Check to the left (unless it's the edge)
	if modulo != 0 {
		if (*hm)[i] >= (*hm)[left_index] {
			return false
		}
	}

	// Check to the right (unless it's the edge)
	if modulo != 99 {
		if (*hm)[i] >= (*hm)[right_index] {
			//fmt.Printf("Right Index %d has value of %d",)
			return false
		}
	}

	// All tests pass, this is a low point!
	fmt.Printf("All tests pass, index %d (%d) is a low point!", i, (*hm)[i])
	return true

}
