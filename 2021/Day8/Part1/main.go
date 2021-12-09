package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Display struct {
	input  []string
	output []string
}

func main() {
	var displays []Display

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "|")

		displays = append(displays, Display{
			input:  strings.Split(line[0], " "),
			output: strings.Split(line[1], " "),
		})

	}

	fmt.Printf("A total of %d displays were found.\n", len(displays))

	count := 0
	for _, v := range displays {
		//go calcFuelUse(v, &fleet_pos, fuel)
		for _, d := range v.output {
			switch len(d) {
			case 2, 3, 4, 7: // 1,7,4,8
				count++
			}
		}
	}

	fmt.Printf("There are %d {1,3,7,8}s\n", count)

}
