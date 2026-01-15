/* A programme that uses a 2D array to store the coordinates of the points in a mathematical relation */

package main

import(
		"fmt"
      )

func main(){
	
	PlottedPoints := [5][2]float64{{34.5, 56.1}, {79.4, 92.9}, {98.8, 61.3}, {34.7, 22.8}, {39.4, 67.1}}
	fmt.Printf("The plotted points are: %0.1f\n", PlottedPoints)
}

