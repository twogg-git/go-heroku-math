package main

import (
	"image/color"

	"github.com/vdobler/chart"
)

func drawBar(xaxis []float64, yaxis []float64) {
	dumper := NewDumper("xbar1", 1, 1, 400, 300)
	defer dumper.Close()

	red := chart.Style{Symbol: 'o', LineColor: color.NRGBA{0xcc, 0x00, 0x00, 0xff},
		FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff},
		LineStyle: chart.SolidLine, LineWidth: 2}
	/*green := chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0xcc, 0x00, 0xff},
	FillColor: color.NRGBA{0x80, 0xff, 0x80, 0xff},
	LineStyle: chart.SolidLine, LineWidth: 2}*/

	barc := chart.BarChart{Title: "Simple Bar Chart"}
	barc.Stacked = false
	barc.Key.Hide = true
	barc.XRange.ShowZero = true
	barc.AddDataPair("Amount",
		xaxis, yaxis, red)
	//[]float64{-10, 10, 20, 30, 35, 40, 50},
	//[]float64{90, 120, 180, 205, 230, 150, 190}, red)
	dumper.Plot(&barc)
	barc.XRange.TicSetting.Delta = 0

	/*
			barc = chart.BarChart{Title: "Simple Bar Chart"}
			barc.Key.Hide = true
			barc.XRange.ShowZero = true
			barc.AddDataPair("Test", []float64{-5, 15, 25, 35, 45, 55}, []float64{110, 80, 95, 80, 120, 140}, green)
			dumper.Plot(&barc)
			barc.XRange.TicSetting.Delta = 0

			barc.YRange.TicSetting.Delta = 0
			barc.Title = "Combined (ugly as bar positions do not match)"
			barc.AddDataPair("Amount", []float64{-10, 10, 20, 30, 35, 40, 50}, []float64{90, 120, 180, 205, 230, 150, 190}, red)
			dumper.Plot(&barc)

			barc.Title = "Stacked (still ugly)"
			barc.Stacked = true
			dumper.Plot(&barc)

			barc = chart.BarChart{Title: "Nicely Stacked"}
			barc.Key.Hide = true
			barc.XRange.Fixed(0, 60, 10)
			barc.AddDataPair("A", []float64{10, 30, 40, 50}, []float64{110, 95, 60, 120}, red)
			barc.AddDataPair("B", []float64{10, 30, 40, 50}, []float64{40, 130, 15, 100}, green)
			dumper.Plot(&barc)

		barc.Stacked = true
		dumper.Plot(&barc)*/
}

/*

func barmain() {
	var debugging = flag.Bool("debug", false, "output debug information to stderr")
	flag.Parse()
	if *debugging {
		chart.DebugLogger = log.New(os.Stdout, "", log.LstdFlags)
	}
	drawBar()
}*/
