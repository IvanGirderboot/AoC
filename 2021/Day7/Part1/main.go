package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fleet_pos []int64
	var fuel = make(chan float64, 1000)
	var lowest_fuel = math.MaxFloat64

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
			pos, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			fleet_pos = append(fleet_pos, pos)
		}

		fmt.Printf("A total of %d crab subs were added to the fleet.\n", len(fleet_pos))
		//fmt.Printf("%v\n", fleet_pos)

		for _, v := range fleet_pos {
			go calcFuelUse(v, &fleet_pos, fuel)
		}

		for i := 0; i < len(fleet_pos); i++ {
			f := <-fuel
			//fmt.Printf("Fuel received in channel: %f\n", f)
			if f < lowest_fuel {
				lowest_fuel = f
			}
		}

		fmt.Printf("The lowest fuel consumption possible is %f\n", lowest_fuel)

	}
}

func calcFuelUse(position int64, fleet *[]int64, fuel chan float64) {
	used_fuel := 0.0

	for i := range *fleet {
		used_fuel += math.Abs(float64((*fleet)[i] - position))
	}
	//fmt.Printf("Aligning to position %d would use %f units of fuel\n", position, used_fuel)
	fuel <- used_fuel
}
