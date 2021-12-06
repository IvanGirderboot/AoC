package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternFish struct {
	timer   int
	created int
}

func main() {
	//var wg sync.WaitGroup
	var school []LanternFish
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
			ticks, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			school = append(school, LanternFish{
				timer:   ticks,
				created: -1})
		}

	}

	fmt.Printf("A total of %d fish were added to the school.\n", len(school))
	for _, v := range school {
		fmt.Printf("%d", v.timer)
	}
	fmt.Printf("\n")

	//wg.Add(len(school))

	for t := 1; t < 81; t++ {
		fmt.Printf("Day %d [%d]: ", t, len(school))

		for i := range school {

			if school[i].timer == 0 {
				school =
					append(school, LanternFish{
						timer:   8,
						created: t,
					})
				school[i].timer = 6
			} else {
				school[i].timer--
			}
		}
		fmt.Printf("There total of %d fish are in the school. after day %d\n", len(school), t)
		/*for _, v := range school {
			fmt.Printf("%d", v.timer)
		}
		fmt.Printf("\n") */

	}
}
