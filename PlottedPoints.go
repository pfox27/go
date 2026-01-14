/* A programme that uses an array to store the coordinates of the set of points plotted 
 * in a mathematical relation. */

package main

import(
		"fmt"
      )

func main(){
	
	var plottedPoints = [5][2]int{{-2,5}, {1,8}, {8,15}, {-5,1}, {-9,-5}}
		for i := 0; i < 5; i++ {
			for j := 0; j < 2; j++ {
				fmt.Printf("plottedPoint[%d][%d] is %d\n" , i, j, plottedPoints[i][j])
				}
		}
}


