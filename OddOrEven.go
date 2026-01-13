/* A programme to accept four numerical arguments and check to see if they are odd or even. */

package main

import(
		"fmt"
		"os"
		"strconv"
      )

func main(){

	val1, _ := strconv.Atoi(os.Args[1])
	val2, _ := strconv.Atoi(os.Args[2])
  	val3, _ := strconv.Atoi(os.Args[3])
	val4, _ := strconv.Atoi(os.Args[4])
	
	fmt.Println("The four numerical values are: ")
	fmt.Printf("%d, %d, %d, %d\n", val1, val2, val3, val4)

	if val1%2 == 0{
		fmt.Println(val1, " is even.")
	} else {fmt.Println(val1, " is odd.")
	}
	if val2%2 == 0{
		fmt.Println(val2, " is even.")
	} else {fmt.Println(val2, " is odd.")
	}
	if val3%2 == 0{
		fmt.Println(val3, " is even.")
	} else {fmt.Println(val3, " is odd.")
	}
	if val4%2 == 0{
		fmt.Println(val4, " is even.")
	} else {fmt.Println(val4, " is odd.")

	}
}
