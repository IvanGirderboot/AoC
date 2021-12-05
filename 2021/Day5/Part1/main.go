package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Line struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func main() {
	var wg sync.WaitGroup
	var lines []Line
	grid := [1000][1000]int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		text = strings.Replace(text, " -> ", ",", 1)
		//fmt.Printf("The points are: %s\n", text)

		split := strings.Split(text, ",")
		var points [4]int
		for i := range split {
			points[i], err = strconv.Atoi(split[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		new_line := Line{
			x1: points[0],
			y1: points[1],
			x2: points[2],
			y2: points[3],
		}
		//fmt.Printf("New line created: X1:%d Y1:%d ---> X2:%d Y2:%d\n", new_line.x1, new_line.y1, new_line.x2, new_line.y2)
		lines = append(lines, new_line)

	}

	fmt.Printf("A total of %d lines were created.\n", len(lines))

	// process all the lines in parrallel
	wg.Add(len(lines))
	for _, v := range lines {
		go addLineToGrid(&grid, v, &wg)
	}
	fmt.Println("Waiting for worker threads")
	wg.Wait()
	fmt.Println("Worker threads complete, calculating results")

	intersecting_points := 0
	for x := range grid {
		//fmt.Printf("ROW %d:%v\n", x, grid[x])
		for y := range grid[x] {
			if grid[x][y] > 1 {
				intersecting_points++
			}
		}
	}

	fmt.Printf("A total of %d points were intersected.\n", intersecting_points)

}

func addLineToGrid(grid *[1000][1000]int, line Line, wg *sync.WaitGroup) {
	defer wg.Done()

	start_index := 0
	end_index := 0
	if line.x1 == line.x2 {
		//fmt.Println("Line is vertical")
		if line.y1 < line.y2 {
			start_index = line.y1
			end_index = line.y2
		} else if line.y1 > line.y2 {
			start_index = line.y2
			end_index = line.y1
		} else {
			log.Fatalf("Line %v is single point", line)
		}
		for i := start_index; i <= end_index; i++ {
			grid[i][line.x1]++
			//fmt.Printf("Marking Point at X:%d Y:%d\n", line.x1, i)
		}
	} else if line.y1 == line.y2 {
		//fmt.Println("Line is horizontal")
		if line.x1 < line.x2 {
			start_index = line.x1
			end_index = line.x2
		} else if line.x1 > line.x2 {
			start_index = line.x2
			end_index = line.x1
		} else {
			log.Fatalf("Line %v is single point", line)
		}
		for i := start_index; i <= end_index; i++ {
			grid[line.y1][i]++
			//fmt.Printf("Marking Point at X:%d Y:%d\n", i, line.y1)
		}

	} else {
		fmt.Printf("Line is at an angle: %v, ignoring\n", line)
	}
}
