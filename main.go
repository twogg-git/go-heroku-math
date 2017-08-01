package main

import (
	"fmt"
	"math/rand"

	"github.com/gonum/stat"
)

func main() {

	// I dont want to do this
	//dataPool := []float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}

	// Instead i want to use rand.Float64() to fill the DataPool, like 100 Float64 numbers

	var dataPool [100]float64
	for i := 0; i < 100; i++ {
		dataPool[i] = rand.Float64()
	}

	// stat.Mean only calculates over []float64 arrays
	mean := stat.Mean(dataPool, nil)
	fmt.Println("Mean value", mean)

	fmt.Println("DataPool...")
	for i := 0; i < len(dataPool); i++ {
		fmt.Println(dataPool[i])
	}
}
