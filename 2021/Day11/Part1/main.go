package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Octopus struct {
	energy_level int64
	flashed      bool
}

func main() {
	var dumbos [100]Octopus
	var flash_count int64

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
			index := (line_number * 10) + i
			dumbos[index] = Octopus{energy_level: num}
		}
		line_number++

	}

	fmt.Printf("A total of %d dumbo octopi were found.\n", len(dumbos))
	displayField(&dumbos)
	//fmt.Printf("%v\n", dumbos)

	for step := 1; step <= 100; step++ {
		// Part 1: Increase all energy levels
		for i := range dumbos {
			dumbos[i].energy_level++
		}

		// Part 2: Flash any octopus with energy level > 9
		for i := range dumbos {
			if dumbos[i].energy_level > 9 {
				flashIfPossible(&dumbos, i)
			}

		}

		// Part 3: count & reset all flashed
		for i := range dumbos {
			if dumbos[i].flashed {
				flash_count++
				dumbos[i].flashed = false
				dumbos[i].energy_level = 0
			}
		}

		fmt.Printf("Step %d: %d Octopi flashes observed thus far.\n", step, flash_count)
		displayField(&dumbos)
		//fmt.Printf("%v\n", dumbos)
	}
}

func flashIfPossible(o *[100]Octopus, i int) bool {
	//fmt.Printf("flashIfPossible called for octopus at position %d\n", i)
	// Increase energy level
	o[i].energy_level++

	// Flash once per round
	if o[i].energy_level > 9 && !o[i].flashed {
		// Set this to flashed
		o[i].flashed = true
	} else {
		return false
		// No need to flash
	}

	up_left_index := i - 11
	up_index := i - 10
	up_right_index := i - 9
	left_index := i - 1
	right_index := i + 1
	down_left_index := i + 9
	down_index := i + 10
	down_right_index := i + 11

	//fmt.Printf("Pos %d Indexes: UL:%d U:%d UR:%d L:%d R:%d DL:%d D:%d DR:%d\n", i, up_left_index, up_index, up_right_index, left_index, right_index, down_left_index, down_index, down_right_index)
	flash_count := 0

	// Check above (unless it's out of bounds)
	if up_index >= 0 {
		//fmt.Println("Flashing Up")
		flash_count++
		flashIfPossible(o, up_index)
	}

	// Check Below (unless it's out of bounds)
	if down_index < len(*o) {
		flash_count++
		//fmt.Println("Flashing Down")
		flashIfPossible(o, down_index)
	}

	modulo := i % 10

	// Check to the left (unless it's the edge)
	if modulo != 0 {
		flash_count++
		//fmt.Println("Flashing Left")
		flashIfPossible(o, left_index)
	}

	// Check to the right (unless it's the edge)
	if modulo != 9 {
		flash_count++
		//fmt.Println("Flashing Right")
		flashIfPossible(o, right_index)
	}

	// Check above-left (unless it's out of bounds)
	if (up_left_index >= 0) && (modulo != 0) {
		flash_count++
		//fmt.Println("Flashing Up-Left")
		flashIfPossible(o, up_left_index)
	}

	// Check above-right (unless it's out of bounds)
	if (up_right_index >= 0) && (modulo != 9) {
		flash_count++
		//fmt.Println("Flashing Up-Right")
		flashIfPossible(o, up_right_index)
	}

	// Check below-left (unless it's out of bounds)
	if (down_left_index < len(*o)) && (modulo != 0) {
		flash_count++
		//fmt.Println("Flashing Down-Left")
		flashIfPossible(o, down_left_index)
	}

	// Check below-right (unless it's out of bounds)
	if (down_right_index < len(*o)) && (modulo != 9) {
		flash_count++
		//fmt.Println("Flashing Down-Right")
		flashIfPossible(o, down_right_index)
	}

	if flash_count != 8 {
		fmt.Printf("===WARNING: Flash count was %d for index %d\n", flash_count, i)
	}
	//fmt.Printf("Index %d has flashed!\n", i)
	return true

}

func displayField(o *[100]Octopus) {
	for i := 0; i < len(o); i++ {
		fmt.Printf("%d ", o[i].energy_level)
		if (i % 10) == 9 {
			fmt.Print("\n")
		}
	}
}
