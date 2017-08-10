package main

import (
	"fmt"
	"math/rand"

	"github.com/gonum/stat"
)

func main() {

	poolSize := 10000
	var dataPool = make([]float64, poolSize)
	for i := 0; i < poolSize; i++ {
		dataPool[i] = rand.Float64()
	}
	
	mean := stat.Mean(dataPool, nil)
	fmt.Println("Mean value", mean)

	stddev := stat.StdDev(dataPool, nil)
	fmt.Println("Standar deviation", stddev)

	/*fmt.Println("DataPool...")
	for i := 0; i < len(dataPool); i++ {
		fmt.Println(dataPool[i])
	}*/
}
