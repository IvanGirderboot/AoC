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

func newDisplay(i []string, o []string) *Display {
	d := Display{
		input:  i,
		output: o,
	}
	d.numbers = make(map[string]string)
	return &d
}

type Display struct {
	input   []string
	output  []string
	numbers map[string]string
}

func (d *Display) sortIO() {
	for i := range d.input {
		//fmt.Printf("formatting input %d: %v to ", i, d.input[i])
		split := strings.Split(strings.Trim(d.input[i], " "), "")
		sort.Strings(split)
		d.input[i] = strings.Join(split, "")
		//fmt.Printf("%s\n", d.input[i])
	}

	// sort inputs by length
	sort.Slice(d.input, func(i, j int) bool {
		return len(d.input[i]) < len(d.input[j])
	})

	for i := range d.output {
		//fmt.Printf("formatting output %d: %v to ", i, d.output[i])
		split := strings.Split(strings.Trim(d.output[i], " "), "")
		sort.Strings(split)
		d.output[i] = strings.Join(split, "")
		//fmt.Printf("%s\n", d.output[i])
	}

}

func (d *Display) decode() {
	d.sortIO()
	//fmt.Printf("Length/Alpha sorted inputs: %v\n", d.input)

	remaining_inputs := d.input

	// some we can tell based on length
	d.numbers[d.input[0]] = "1"
	one := d.input[0]
	//fmt.Printf("One is: %s \n", one)

	d.numbers[d.input[1]] = "7"
	seven := d.input[1]
	//fmt.Printf("Seven is: %s \n", seven)

	d.numbers[d.input[2]] = "4"
	four := d.input[2]
	//fmt.Printf("Four is: %s \n", four)

	d.numbers[d.input[9]] = "8"
	//eight := d.input[9]
	//fmt.Printf("Eight is: %s \n", eight)

	remaining_inputs = remove(remaining_inputs, 9)
	remaining_inputs = remove(remaining_inputs, 2)
	remaining_inputs = remove(remaining_inputs, 1)
	remaining_inputs = remove(remaining_inputs, 0)

	//fmt.Printf("Unsolved Inputs: %v\n", remaining_inputs)

	// Nine
	//nine := ""
	for i, digit := range remaining_inputs {
		//fmt.Printf("Testing for nine in combo: %v\n", digit)
		if len(digit) != 6 {
			continue
		}
		match_count := 0

		for _, wire := range digit {

			if strings.ContainsRune(four, wire) || strings.ContainsRune(seven, wire) {
				match_count++
			}
		}
		if match_count == 5 {
			d.numbers[digit] = "9"
			remaining_inputs = remove(remaining_inputs, i)
			//fmt.Printf("Unsolved Inputs: %v\n", remaining_inputs)
			//nine = digit
			//fmt.Printf("Nine is %v\n", nine)
			break
		}
	}

	// Zero and Six
	//zero := ""
	//six := ""
	for _, digit := range remaining_inputs {
		if len(digit) != 6 {
			continue
		}
		match_count := 0

		for _, wire := range digit {

			if strings.ContainsRune(seven, wire) {
				match_count++
			}
		}
		if match_count == 3 {
			//zero = digit
			//fmt.Printf("Zero is %s\n", zero)
			d.numbers[digit] = "0"
		} else if match_count == 2 {
			//six = digit
			//fmt.Printf("Six is %s\n", six)
			d.numbers[digit] = "6"
		} else {
			log.Fatalf("Error!")
		}
	}
	remaining_inputs = remaining_inputs[0:3]
	//fmt.Printf("Unsolved Inputs: %v\n", remaining_inputs)

	// Three
	three := ""
	for i, digit := range remaining_inputs {
		//fmt.Printf("Testing for three in combo: %v\n", digit)

		match_count := 0

		for _, wire := range digit {

			if strings.ContainsRune(one, wire) {
				match_count++
			}
		}
		if match_count == 2 {
			three = digit
			//fmt.Printf("Three is %v\n", three)
			d.numbers[three] = "3"
			remaining_inputs = remove(remaining_inputs, i)
			//fmt.Printf("Unsolved Inputs: %v\n", remaining_inputs)
			break
		}
	}
	// Two & Five
	// Zero and Six
	//two := ""
	//five := ""
	for _, digit := range remaining_inputs {
		match_count := 0

		for _, wire := range digit {
			if strings.ContainsRune(four, wire) {
				match_count++
			}
		}
		if match_count == 3 {
			//five = digit
			//fmt.Printf("Five is %s\n", five)
			d.numbers[digit] = "5"
		} else if match_count == 2 {
			//two = digit
			//fmt.Printf("Two is %s\n", two)
			d.numbers[digit] = "2"
		} else {
			log.Fatalf("Error!")
		}
	}

	//fmt.Printf("Known Numbers: %v\n", d.numbers)

}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	var displays []Display
	var total int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "|")

		line[0] = strings.Trim(line[0], " ")
		line[1] = strings.Trim(line[1], " ")

		displays = append(displays, *newDisplay(strings.Split(line[0], " "), strings.Split(line[1], " ")))

	}

	fmt.Printf("A total of %d displays were found.\n", len(displays))

	count := 0
	for i, v := range displays {
		displays[i].decode()

		output_string := ""
		for _, d := range v.output {
			switch len(d) {
			case 2, 3, 4, 7: // 1,7,4,8
				count++
			}
			num, found := displays[i].numbers[d]
			if found {
				output_string += num
			} else {
				log.Fatalf("Unable to decode output digit %v", d)
			}
		}

		number, err := strconv.ParseInt(output_string, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total += number
		fmt.Printf("Display %d has an output of %d\n", i, number)
	}

	fmt.Printf("There are %d {1,3,7,8}s\n", count)
	fmt.Printf("∑ outputs: %d \n", total)
}
