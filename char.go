package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gonum/stat"
)

func main() {
	fmt.Println("Starting server...")
	poolSize := 250
	drawBar(randomXAxis(poolSize), initPool(poolSize))
	http.HandleFunc("/", calltemplate)
	http.ListenAndServe(":8080", nil)
}

func calltemplate(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error stating server!", err)
	}
}

func randomXAxis(poolSize int) []float64 {

	const randLimit int = 10
	fmt.Println("Random limit", randLimit)

	var dataPool = make([]float64, poolSize)
	for i := 0; i < poolSize; i++ {
		dataPool[i] = float64(rand.Intn(randLimit)) + rand.Float64()
	}

	mean := stat.Mean(dataPool, nil)
	fmt.Println("Mean value", mean)

	stddev := stat.StdDev(dataPool, nil)
	fmt.Println("Standar deviation", stddev)

	return dataPool
}

func initPool(poolSize int) []float64 {

	var dataPool = make([]float64, poolSize)
	for i := 0; i < poolSize; i++ {
		dataPool[i] = float64(i)
	}

	return dataPool
}

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>Image sample from GitHub</h2>
	<img 
		src="http://raw.githubusercontent.com/twogg-git/go-math/master/xbar1.png" 
		alt="chart" 
		style="width:600px;height:300px;"
	>
</center></body></html>
`))
