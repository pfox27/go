/* A programme that uses an array to store rainfall stats and a loop to find the total rainfall. */

package main

import(
		"fmt"
      )

func main(){

	totalRainfall := 0.0
	meanRainfall := 0.0
	
	monthlyRainfall := [12]float64{34.56, 56.12, 79.45, 92.96, 98.83, 61.33, 34.70, 22.88, 39.42, 67.18, 71.22, 43.99}
		for _, monthlyValue := range monthlyRainfall{
	       totalRainfall += monthlyValue
		}
	
	meanRainfall = totalRainfall/float64(len(monthlyRainfall))
	fmt.Printf("The monthly rainfall in mm is %0.1f\n", monthlyRainfall)
	fmt.Printf("The total rainfall for the %d months entered is %0.1f mm.\n", len(monthlyRainfall), totalRainfall)
	fmt.Printf("The mean rainfall for the %d months entered is %0.1f mm.\n", len(monthlyRainfall), meanRainfall)

}

