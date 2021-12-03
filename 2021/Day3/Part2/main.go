package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var text_data []string

	var co2_scrubber_rate int64
	var o2_gen_rate int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		text_data = append(text_data, text)
	}

	o2_gen_rate_choices := text_data
	for bit := 0; bit < 12; bit++ {
		val := findMostCommonBitInPos(o2_gen_rate_choices, bit, "1")
		o2_gen_rate_choices = filterSliceByBit(o2_gen_rate_choices, bit, val)
		fmt.Printf("There are %d possible values after bit %d\n", len(o2_gen_rate_choices), bit+1)
		if len(o2_gen_rate_choices) == 1 {
			o2_gen_rate, err = strconv.ParseInt(o2_gen_rate_choices[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	co2_scrubber_rate_choices := text_data
	for bit := 0; bit < 12; bit++ {
		val := findLeastCommonBitInPos(co2_scrubber_rate_choices, bit, "0")
		co2_scrubber_rate_choices = filterSliceByBit(co2_scrubber_rate_choices, bit, val)
		fmt.Printf("There are %d possible values after bit %d\n", len(co2_scrubber_rate_choices), bit+1)
		if len(co2_scrubber_rate_choices) == 1 {
			co2_scrubber_rate, err = strconv.ParseInt(co2_scrubber_rate_choices[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	fmt.Printf("O2 Scrubber rate is %d and CO2 Scrubber rate is %d with a product of %d", o2_gen_rate, co2_scrubber_rate, o2_gen_rate*co2_scrubber_rate)
}

func filterSliceByBit(input []string, bit_pos int, bit_val string) []string {
	var filtered []string
	for _, reading := range input {
		if reading[bit_pos:bit_pos+1] == bit_val {
			filtered = append(filtered, reading)
		}
	}
	return filtered
}

func findMostCommonBitInPos(input []string, bit_pos int, tie_breaker_value string) string {
	var zero int
	var one int
	for _, v := range input {
		if v[bit_pos:bit_pos+1] == "0" {
			zero++
		} else {
			one++
		}
	}
	if zero > one {
		return "0"
	} else if zero < one {
		return "1"
	}
	fmt.Println("We have a tie!")
	return tie_breaker_value

}

func findLeastCommonBitInPos(input []string, bit_pos int, tie_breaker_value string) string {
	var zero int
	var one int
	for _, v := range input {
		if v[bit_pos:bit_pos+1] == "0" {
			zero++
		} else {
			one++
		}
	}
	if zero < one {
		return "0"
	} else if zero > one {
		return "1"
	}
	fmt.Println("We have a tie!")
	return tie_breaker_value

}
