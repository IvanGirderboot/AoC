package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var score int64
	var line_number int
	var completion_scores []int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_number++
		//fmt.Printf("===========Processing line number %d================\n", line_number)
		index, err_char, end_line := detectClosingElement(line, 0)
		if err_char != "" {
			fmt.Printf("Line %d was corrupted at index %d with character %s\n", line_number, index, err_char)

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

		}
		if line != end_line {
			added_chars := end_line[len(line):]
			fmt.Printf("Line %d was incomplete, added %d characters: %s\n", line_number, len(end_line)-len(line), added_chars)
			completion_score := 0
			for _, c := range added_chars {
				completion_score = completion_score * 5
				switch string(c) {
				case ")":
					completion_score += 1
				case "]":
					completion_score += 2
				case "}":
					completion_score += 3
				case ">":
					completion_score += 4
				default:
					log.Fatalf("Unexpected completion character %s\n", string(c))
				}
				//fmt.Printf("Completion score is now %d\n", completion_score)
			}
			//fmt.Printf("Completion score total for line %d is %d\n", line_number, completion_score)
			completion_scores = append(completion_scores, int64(completion_score))

		}
		//fmt.Printf("Done processing line at index %d\n", index)

	}
	fmt.Printf("Final error score is %d\n", score)
	//fmt.Printf("There are %d Completion Scores\n", len(completion_scores))
	sort.Slice(completion_scores, func(i, j int) bool { return completion_scores[i] > completion_scores[j] })
	median := len(completion_scores) / 2
	fmt.Printf("Median completion score (Index %d) is %d\n", median, completion_scores[median])

}

// Given a string and index of an opening character, find the right closing character
func detectClosingElement(line string, i int) (index int, unexpected_char string, new_line string) {
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
	//fmt.Printf("Looking for closing character %s\n", closing_character)

	const opening_characters = "{[(<"

	index++
	//fmt.Printf("Increment: Index is %d\n", index)

	// If we see another opening charater, recurse!
check:
	if index > len(line)-1 { //Detect incomplete lines
		//fmt.Printf("Found incomplete line at index %d\n", index)
		//return index, ""
		// Append missing character
		//fmt.Printf("Added character %s", closing_character)
		line += closing_character
	}

	if strings.ContainsAny(line[index:index+1], opening_characters) {
		//fmt.Printf("Detected opening character %s at index %d, diving deeper.\n", line[index:index+1], index)
		new_index, u, l := detectClosingElement(line, index)
		if u != "" { // We can stop processing and return the corrupted character right away
			//fmt.Printf("Returned error index is %d\n", new_index)
			return new_index, u, l
		}
		//fmt.Printf("Recurse: Index is %d\n", new_index)
		index = new_index
		line = l
		goto check // Start the check over
	}

	//fmt.Printf("Check: Index is %d\n", index)
	if line[index:index+1] == closing_character {
		//fmt.Printf("Yay! Closing character %s detected at index %d!\n", closing_character, index)
		return index + 1, "", line
	}

	//fmt.Printf("Error! Incorrect closing character detected at index %d!  Expected %s, got %s\n", index, closing_character, line[index:index+1])
	return index, line[index : index+1], line
}
