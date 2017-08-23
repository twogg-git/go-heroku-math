package main

import (
	"fmt"
	"math/rand"

	"net/http"

	"github.com/gonum/stat"
)

func main() {
	http.HandleFunc("/", statTest)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func statTest(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "Testing , gonum/stat fuctions")

	poolSize := 100000
	fmt.Fprintln(res, "Pool size", poolSize)

	var dataPool = make([]float64, poolSize)
	for i := 0; i < poolSize; i++ {
		dataPool[i] = rand.Float64()
	}

	// stat.Mean only calculates over []float64 arrays
	mean := stat.Mean(dataPool, nil)
	fmt.Fprintln(res, "Mean value", mean)

	stddev := stat.StdDev(dataPool, nil)
	fmt.Fprintln(res, "Standar deviation", stddev)

	/*fmt.Println("DataPool...")
	for i := 0; i < len(dataPool); i++ {
		fmt.Println(dataPool[i])
	}*/
}
