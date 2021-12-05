package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GameBoard struct {
	numbers []string // don't define the board size here becuse using append() is easier
	marked  [25]bool
}

func (gb *GameBoard) markBoard(input string) (wasMarked bool) {
	for i, v := range gb.numbers {
		if v == input {
			gb.marked[i] = true
			return true
		}
	}
	return false // Number was not on board
}

func (gb *GameBoard) checkIfWinner(winning_values []string) (isWinner bool, details string) {
	//check the rows
	if checkLineForMarks(gb.marked[0:5]) {
		return true, "row 1"
	}

	if checkLineForMarks(gb.marked[5:10]) {
		return true, "row 2"
	}

	if checkLineForMarks(gb.marked[10:15]) {
		return true, "row 3"
	}

	if checkLineForMarks(gb.marked[15:20]) {
		return true, "row 4"
	}

	if checkLineForMarks(gb.marked[20:25]) {
		return true, "row 5"
	}

	// check the columsn
	col1 := []bool{gb.marked[0], gb.marked[5], gb.marked[10], gb.marked[15], gb.marked[20]}
	if checkLineForMarks(col1) {
		return true, "column 1"
	}

	col2 := []bool{gb.marked[1], gb.marked[6], gb.marked[11], gb.marked[16], gb.marked[21]}
	if checkLineForMarks(col2) {
		return true, "column 2"
	}

	col3 := []bool{gb.marked[2], gb.marked[7], gb.marked[12], gb.marked[17], gb.marked[22]}
	if checkLineForMarks(col3) {
		return true, "column 3"
	}

	col4 := []bool{gb.marked[3], gb.marked[8], gb.marked[13], gb.marked[18], gb.marked[23]}
	if checkLineForMarks(col4) {
		return true, "column 4"
	}

	col5 := []bool{gb.marked[4], gb.marked[9], gb.marked[14], gb.marked[19], gb.marked[24]}
	if checkLineForMarks(col5) {
		return true, "column 5"
	}

	// check the diagnals
	diag1 := []bool{gb.marked[0], gb.marked[6], gb.marked[12], gb.marked[18], gb.marked[24]}
	if checkLineForMarks(diag1) {
		return true, "upper left to lower right diagnal 1"
	}

	diag2 := []bool{gb.marked[4], gb.marked[8], gb.marked[12], gb.marked[16], gb.marked[20]}
	if checkLineForMarks(diag2) {
		return true, "lower left to upper right diagnal 1"
	}

	return false, ""
}

func (gb *GameBoard) unmarkedSum() (sum int) {
	for i, v := range gb.marked {
		if !v {
			num, err := strconv.Atoi(gb.numbers[i])
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}
	}
	return sum
}

func addBoardData(input []string, boards *[]GameBoard) {
	i := len(*boards) - 1
	if (i < 0) || (len((*boards)[i].numbers) >= 25) {
		i++
		new_board := GameBoard{}
		*boards = append(*boards, new_board)
	}
	fmt.Printf("Updating board at index %d with numbers %v\n", i, input)
	(*boards)[i].numbers = append((*boards)[i].numbers, input...)

}

func checkLineForMarks(line []bool) bool {
	for i, v := range line {

		if !v {
			break
		}
		if i == 4 {
			return true
		}
	}
	return false
}

func main() {
	var called_nums []string
	var boards []GameBoard
	//var winning_boards []GameBoard

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Called Numbers
		if strings.Contains(text, ",") {
			called_nums = append(called_nums, strings.Split(text, ",")...)
			fmt.Printf("The called numbers are: %v\n", called_nums)
			continue
		}

		// Board Rows
		r := regexp.MustCompile(`(\d{1,2})`)
		digit_matches := r.FindAllString(text, -1)
		if digit_matches != nil {
			addBoardData(digit_matches, &boards)
			continue
		}

		fmt.Printf("Skipping row with data '%s'\n", text)
		continue
	}

	fmt.Printf("A total of %d board were populated.\n", len(boards))

	// Sanity check all boards
	for i, v := range boards {
		if len(v.numbers) != 25 {
			log.Fatalf("Board at index %d has %d numbers, expecte to see 25", i, len(v.numbers))
		}
	}
	fmt.Println("All boards have 25 numbers. (That's a good thing!)")

number_calling:
	for i, v := range called_nums {
		current_called_nums := called_nums[:i+1]

		//fmt.Printf("Calling number %d: %s\n", i+1, v)
		fmt.Printf("Called numbers thus far: %s\n", current_called_nums)

		for bi := range boards {
			mark := boards[bi].markBoard(v)

			if mark {
				//fmt.Printf("Board index %d was marked with %s\n", bi, v)
				if len(current_called_nums) < 5 {
					fmt.Println("No need to check for winners with less than 5 numbers called.")
					continue
				}

				winner, deets := boards[bi].checkIfWinner(current_called_nums)
				if winner {
					fmt.Printf("Board index %d is a winner on %s!!!111one\n", bi, deets)

					last_number, _ := strconv.Atoi(v)
					fmt.Printf("Winning board's unmarked sum is: %d, last called number was %d with score of %d\n", boards[bi].unmarkedSum(), last_number, boards[bi].unmarkedSum()*last_number)
					break number_calling
				}
			}
		}
	}
}
