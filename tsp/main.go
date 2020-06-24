package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//City has City Coordinates
type City struct {
	x float64
	y float64
}

func (a *City) DistanceTo(b City) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)

	fd := math.Sqrt((dx * dx) + (dy * dy))
	return fd
}

//readInput is reads a csv file and returns a slice of the city's coordinates.
func readInput(in string) []City {
	//Path of the file to be read
	path := "../input_" + in + ".csv"

	cities := make([]City, 0)
	//Load a file
	file, err := os.Open(path)
	//If an error occurs, display an error and exit.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read one line at a time and repeat.
	for scanner.Scan() {
		line := scanner.Text()
		xy := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(xy[0], 32)
		y, _ := strconv.ParseFloat(xy[1], 32)
		cities = append(cities, City{x, y})
	}
	return cities
}

//printSolution writes the result to a csv file.
func printSolution(in string, solution []string) {
	path := "../solution_yours_" + in + ".csv"
	//書き込み用に開く
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.Write(solution); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	writer.Flush()
}

func main() {
	var in = flag.String("input", "0", "input number is 0-6")
	flag.Parse()
	cities := readInput(*in)
	var solution []string

	for i := 1; i < len(cities); i++ {
		fmt.Println(len(cities))
		if i == len(cities)-1 {
			s := strconv.Itoa(int(cities[i].DistanceTo(cities[1])))
			solution = append(solution, s)
			continue
		}
		s := strconv.Itoa(int(cities[i].DistanceTo(cities[i+1])))
		solution = append(solution, s)
	}

	printSolution(*in, solution)

}
