package main

import (
	"fmt"
	"html/template"
	"image/color"
	"math/rand"
	"net/http"
	"time"

	"github.com/gonum/stat"
	"github.com/vdobler/chart"
)

const poolSize int = 250

func chartmain() {
	fmt.Println("Starting server...")

	drawBar(randomXAxis(poolSize), initPool(poolSize))
	http.HandleFunc("/", calltemplate)
	http.ListenAndServe(":3001", nil)
}

func calltemplate(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error stating server!", err)
	}
}

func randomXAxis(poolSize int) []float64 {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Random limit", randInt(0, 100))

	randLimit := rand.Intn(100)
	fmt.Println("Random limit", randLimit)

	var dataPool = make([]float64, poolSize)
	for i := 0; i < poolSize; i++ {
		dataPool[i] = float64(rand.Intn(randLimit)) + rand.Float64()
	}

	totalSum := 0.0
	for _, number := range dataPool {
		totalSum = totalSum + number
	}
	mean2 := totalSum / float64(poolSize)
	fmt.Println("Mean manual", mean2)

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

func drawBar(xaxis []float64, yaxis []float64) {
	dumper := NewDumper("xbar1", 1, 1, 400, 300)
	defer dumper.Close()

	red := chart.Style{Symbol: 'o', LineColor: color.NRGBA{0xcc, 0x00, 0x00, 0xff},
		FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff},
		LineStyle: chart.SolidLine, LineWidth: 2}

	barc := chart.BarChart{Title: "Histogram Ramdon Numbers"}
	barc.Stacked = false
	barc.Key.Hide = true

	barc.XRange.ShowZero = false
	barc.XRange.Fixed(0, 100, 10)
	barc.YRange.Fixed(0, 250, 25)
	barc.AddDataPair("Datapool", xaxis, yaxis, red)

	dumper.Plot(&barc)
	barc.XRange.TicSetting.Delta = 0

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
	<h1>Histogram Ramdon Numbers</h1>
	<h2>Data pool size : 250 numbers</h2>
	<h2>Data pool range : 0-100 </h2>
	<img 
		src="file:///Users/Cruz/go/src/go-math/xbar1.png" 
		alt="local file" 
		style="width:600px;height:300px;"
	>
	<img 
		src="http://raw.githubusercontent.com/twogg-git/go-math/master/xbar1.png" 
		alt="chart" 
		style="width:600px;height:300px;"
	>
</center></body></html>
`))
