package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type City struct {
	x float64
	y float64
}

func readInput(filename string) []City {
	cities := make([]City, 0)
	//Load a file
	file, err := os.Open(filename)
	//If an error occurs, display an error and exit.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Load from file
	scanner := bufio.NewScanner(file)
	// Read one line at a time and repeat.
	for scanner.Scan() {
		xy := strings.Split(scanner.Text(), ",")
		x, err := strconv.ParseFloat(xy[0], 32)
		y, err := strconv.ParseFloat(xy[1], 32)
		cities = append(cities, City{x, y})
	}
	return cities
}
