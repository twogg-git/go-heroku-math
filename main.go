package main

import (
	"fmt"
	"net/http"
	"os"
  	"time"
	"math/rand"
  	"github.com/gonum/stat"
)

func main() {
	http.HandleFunc("/", mathTest)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func mathTest(res http.ResponseWriter, req *http.Request) {
	
  	fmt.Fprintln(res, "Hello, datapool testing")
  
  	now := time.Now()
  	fmt.Fprintln(res, now)
  	
  	const poolSize int = 10000
  	const randLimit int = 2500
  	var dataPool2 [poolSize]int
  
  	fmt.Fprintln(res, "Adding data into the pool")
  	for i:= 0; i < poolSize; i++ {
      	dataPool2[i] = rand.Intn(randLimit)
  	}
  	
  	fmt.Fprintln(res, "Pool size")
  	fmt.Fprintln(res, poolSize)
    fmt.Fprintln(res, "Random number limit")             
  	fmt.Fprintln(res, randLimit)

  	fmt.Fprintln(res, "Datapool size", len(dataPool2))
  	
  	media := stat.Mean(dataPool2)	
  	fmt.Fprintln(res, "Median")
  	fmt.Fprintln(res, media)
  
  	for i:=0; i < len(dataPool2); i++ {
      	fmt.Fprintln(res, dataPool2[i])
  	} 
  
 	fmt.Fprintln(res, "Yes it's alive!!!")

}
