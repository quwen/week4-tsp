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

type Tour struct {
	tourCities []City
	distance   float64
}

func (a *City) DistanceTo(b City) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)

	fd := math.Sqrt((dx * dx) + (dy * dy))
	return fd
}

func (a *Tour) GetCity(tourPosition int) City {
	return a.tourCities[tourPosition]
}

func (a *Tour) TourSize() int {
	return len(a.tourCities)
}

func (a *Tour) TourDistance() float64 {
	if a.distance == 0 {
		td := float64(0)
		for i := 0; i < a.TourSize(); i++ {
			fromC := a.GetCity(i)
			var destC City
			if i+1 < a.TourSize() {
				destC = a.GetCity(i + 1)
			} else {
				destC = a.GetCity(0)
			}
			td += fromC.DistanceTo(destC)
		}
		a.distance = td
	}
	return a.distance
}

func min(a []float64) float64 {
	min := 0
	for i, v := range a {
		if v < a[min] {
			min = i
		}
	}
	return min
}

func solve(cities []City) int {
	var solution []int
	solution = append(solution, 0)

	var dst [][]float64

	for i := 1; i < len(cities); i++ {
		for j := 1; j < len(cities); i++ {
			if i == j;continue
		dst[i-1] = append(dst[i-1], cities[i].DistanceTo(cities[j]))
	}

	var unvcities []City
	for i := 1; i < len(cities); i++ {
	unvcities = append(unvcities, cities[i])
	}
	city := 0
	for i := 0; i < len(unvities); {
		ncity := min(unvcities(dist[city]))
	}




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
func printSolution(in string, solution [][]string) {
	path := "../solution_yours_" + in + ".csv"
	//書き込み用に開く
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	for _, s := range solution {
		if err := writer.Write(s); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	writer.Flush()
}

func main() {
	var in = flag.String("input", "0", "input number is 0-6")
	flag.Parse()
	cities := readInput(*in)

	var solution [][]string
	solution = append(solution, []string{"index"})

	for i := 1; i < len(cities); i++ {
		fmt.Println(len(cities))
		if i == len(cities)-1 {
			s := strconv.Itoa(int(cities[i].DistanceTo(cities[1])))
			solution = append(solution, []string{s})
			continue
		}
		s := strconv.Itoa(int(cities[i].DistanceTo(cities[i+1])))
		solution = append(solution, []string{s})
	}

	printSolution(*in, solution)

}
