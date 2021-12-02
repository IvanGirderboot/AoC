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

	var horiz int64 = 0
	var depth int64 = 0

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), " ")
		value, err := strconv.ParseInt(d[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		switch d[0] {
		case "up":
			depth = depth - value
		case "down":
			depth = depth + value
		case "forward":
			horiz = horiz + value
		default:
			log.Fatalf("Cound not parse input %s", scanner.Text())
		}

		fmt.Printf("Current instruction is %s by %d | horizontal:%d depth:%d\n", d[0], value, horiz, depth)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Final Position is horizontal:%d depth:%d with a product of %d", horiz, depth, horiz*depth)
}
