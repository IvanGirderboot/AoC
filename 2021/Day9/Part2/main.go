package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var height_map [10000]int64
	var basin_map [10000]bool
	var risk_level_total int64
	var basin_sizes []int64

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
			fmt.Printf("Calculating Basin Map for point %d\n", i)
			basin_sizes = append(basin_sizes, discoverBasin(&height_map, &basin_map, int64(i)))
		}
	}

	fmt.Printf("Risk level Total is %d\n", risk_level_total)

	sort.Slice(basin_sizes, func(i, j int) bool { return basin_sizes[i] > basin_sizes[j] })
	fmt.Printf("Basins discovered: %v\n", basin_sizes)

	product := basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
	fmt.Printf("Product of top 3 basins is: %d\n", product)
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
	fmt.Printf("All tests pass, index %d (%d) is a low point!\n", i, (*hm)[i])
	return true

}

// Discover basin takes a starting point and checks it's neighbors to see if they
//   are also part of the basin (recursively calling itself on said neighbors)
func discoverBasin(hm *[10000]int64, bm *[10000]bool, i int64) (size int64) {
	if (*bm)[i] {
		fmt.Printf("Basin point at %d is already explored\n", i)
		return 0
	}

	up_index := i - 100
	down_index := i + 100
	left_index := i - 1
	right_index := i + 1

	size++          // Increment for this point
	(*bm)[i] = true // Mark this point as already explored

	// Check above (unless it's out of bounds)
	if up_index > 0 {
		if ((*hm)[up_index] > (*hm)[i]) && ((*hm)[up_index] < 9) {
			size += discoverBasin(hm, bm, up_index)
		}
	}

	// Check Below (unless it's out of bounds)
	if down_index < int64(len(*hm)-1) {
		if ((*hm)[down_index] > (*hm)[i]) && ((*hm)[down_index] < 9) {
			size += discoverBasin(hm, bm, down_index)
		}
	}

	modulo := i % 100

	// Check to the left (unless it's the edge)
	if modulo != 0 {
		if ((*hm)[left_index] > (*hm)[i]) && ((*hm)[left_index] < 9) {
			size += discoverBasin(hm, bm, left_index)
		}
	}

	// Check to the right (unless it's the edge)
	if modulo != 99 {
		if ((*hm)[right_index] > (*hm)[i]) && ((*hm)[right_index] < 9) {
			//fmt.Printf("Right Index %d has value of %d",)
			size += discoverBasin(hm, bm, right_index)
		}
	}

	// We made it!  return the size back up the stack!
	fmt.Printf("Basin size from point %d is %d\n", i, size)

	return size

}
