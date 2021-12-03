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

	var gam []string
	var epi []string
	var text_data []string

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

	for bit := 0; bit < 12; bit++ {
		gam = append(gam, findMostCommonBitInPos(text_data, bit, "1"))
		epi = append(epi, findLeastCommonBitInPos(text_data, bit, "1"))
	}

	fmt.Printf("Gamma is: %s\n", strings.Join(gam, ""))
	fmt.Printf("Episilon is:%s\n", strings.Join(epi, ""))

	gam_joined := strings.Join(gam, "")
	gam_decimal, err := strconv.ParseInt(gam_joined, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epi_joined := strings.Join(epi, "")
	epi_decimal, err := strconv.ParseInt(epi_joined, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gam is %d, epi is %d in decimal\n", gam_decimal, epi_decimal)
	fmt.Printf("power output is %d\n", (gam_decimal * epi_decimal))

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
