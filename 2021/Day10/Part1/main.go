package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var score int64
	var line_number int

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_number++
		fmt.Printf("===========Processing line number %d================\n", line_number)
		index, err_char := detectClosingElement(line, 0)
		if err_char != "" {
			fmt.Printf("Error at index %d with character %s\n", index, err_char)

			switch err_char {
			case ")":
				score += 3
			case "]":
				score += 57
			case "}":
				score += 1197
			case ">":
				score += 25137
			default:
				log.Fatalf("Unknown illegal character %s\n", err_char)
			}

			fmt.Printf("Current error score is %d\n", score)
		}
		fmt.Printf("Done processing line at index %d\n", index)

	}
	fmt.Printf("Final error score is %d", score)
}

// Given a string and index of an opening character, find the right closing character
func detectClosingElement(line string, i int) (index int, unexpected_char string) {
	index = i
	//fmt.Printf("Start: Index is %d\n", index)
	closing_character := ""
	switch string(line[i]) {
	case "{":
		closing_character = "}"
	case "[":
		closing_character = "]"
	case "(":
		closing_character = ")"
	case "<":
		closing_character = ">"
	default:
		log.Fatalf("Unexpected opening character %s\n", line[index:index+1])
	}
	fmt.Printf("Looking for closing character %s\n", closing_character)

	const opening_characters = "{[(<"

	index++
	//fmt.Printf("Increment: Index is %d\n", index)

	// If we see another opening charater, recurse!
check:
	if index > len(line)-1 { //Ignore incomplete lines
		fmt.Printf("Ignoring incomplete line at index %d\n", index)
		return index, ""
	}

	if strings.ContainsAny(line[index:index+1], opening_characters) {
		fmt.Printf("Detected opening character %s at index %d, diving deeper.\n", line[index:index+1], index)
		new_index, u := detectClosingElement(line, index)
		if u != "" { // We can stop processing and return the corrupted character right away
			fmt.Printf("Returned error index is %d\n", new_index)
			return new_index, u
		}
		//fmt.Printf("Recurse: Index is %d\n", new_index)
		index = new_index
		goto check // Start the check over
	}

	//fmt.Printf("Check: Index is %d\n", index)
	if line[index:index+1] == closing_character {
		fmt.Printf("Yay! Closing character %s detected at index %d!\n", closing_character, index)
		return index + 1, ""
	}

	fmt.Printf("Error! Incorrect closing character detected at index %d!  Expected %s, got %s\n", index, closing_character, line[index:index+1])
	return index, line[index : index+1]
}
