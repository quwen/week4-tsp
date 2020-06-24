package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type City struct {
	x float64
	y float64
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

	// Read one line at a time and repeat.
	for scanner.Scan() {
		line := scanner.Text()
		xy := strings.Split(line, ",")
		x, err := strconv.ParseFloat(xy[0], 32)
		y, err := strconv.ParseFloat(xy[1], 32)
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
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	for _, s := range solution {
		if err := writer.Write(s); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
}

func main() {
	var n = flag.String("input number", "0", "input number is 0-6")
	flag.Parse()
	cities := readInput(*n)

	var solution []string

	printSolution(*in, solution)

}
