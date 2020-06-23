package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type City struct {
	x float64
	y float64
}

func readInput(num string) []City {
	//Path of the file to be read
	path := "../input_" + num + ".csv"

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
