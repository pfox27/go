/* A programme that uses an array to store the coordinates of the set of points plotted 
 * in a mathematical relation. */

package main

import(
		"fmt"
      )

func main(){
	
	var groceryList = [3]string{"bananas", "eggs", "hummus"}
	fmt.Printf("Original List: %s\n" , groceryList)
	changeList(groceryList)
	
				}
		}
}


func changeList(arr [3] string){
	
	//Change an item in the list
	
	arr[3] = "queso dip"
	
