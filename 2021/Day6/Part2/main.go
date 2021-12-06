package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {

	var school [9]big.Int

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		for _, v := range split {
			age, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			school[age].Add(&school[age], big.NewInt(int64(1)))
		}

	}
	fmt.Printf("Start :")
	for i, v := range school {
		fmt.Printf("Age:%d Count:%d | ", i, v.Uint64())
	}
	fmt.Printf("\n")

	for t := 0; t < 257; t++ {
		fmt.Printf("Day %d :", t)
		total := big.NewInt(int64(0))

		new_fish := school[0] // should go to timer 8

		for i := 0; i <= 7; i++ {
			// shift fish down a day
			school[i] = school[i+1]

			// Day 6 gets all the former day 0 fish added to it
			if i == 6 {
				school[i].Add(&school[i], &new_fish)
			}

			total.Add(total, &school[i])
			fmt.Printf("Age:%d Count:%d | ", i, school[i].Uint64())
		}

		// Last day get's all the new fish
		school[8] = new_fish
		fmt.Printf("Age:%d Count:%d | ", 8, school[8].Uint64())
		fmt.Printf("TOTAL:%d\n", total)
	}
}
