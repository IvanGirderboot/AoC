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
	var lastNumber int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		currentNumber, err := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Printf("Current number is: %d -- ", currentNumber)

		if err != nil {
			log.Fatal(err)
		}

		// Don't calculate on first number
		if lastNumber == 0 {
			lastNumber = currentNumber
			continue
		}

		if currentNumber > lastNumber {
			increases++
			fmt.Println("Increased!")
		} else {
			decreases++
			fmt.Println("Decreased!")
		}
		lastNumber = currentNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Totals: Increases: %d Decreases: %d", increases, decreases)
}
