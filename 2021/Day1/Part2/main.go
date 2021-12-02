package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var increases int64
	var decreases int64
	var numbers []int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		currentNumber, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Current number is: %d; ", currentNumber)

		numbers = append(numbers, currentNumber)

		// Don't calculate anything until we have at least 4 elements
		if len(numbers) < 4 {
			fmt.Println("Skipping until we have more data!")
			continue
		}

		last_number_index := len(numbers) - 1

		current_3pt_sum := numbers[last_number_index] + numbers[last_number_index-1] + numbers[last_number_index-2]
		previous_3pt_sum := numbers[last_number_index-1] + numbers[last_number_index-2] + numbers[last_number_index-3]
		fmt.Printf("Current 3-Point total is: %d, Previous 3-Point total is %d -- ", current_3pt_sum, previous_3pt_sum)

		if current_3pt_sum > previous_3pt_sum {
			increases++
			fmt.Println("Increased!")
		} else {
			decreases++
			fmt.Println("Decreased!")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Totals: Increases: %d Decreases: %d", increases, decreases)
}
