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
	fmt.Println("\n")
		for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s is %s\n", os.Args[i], oddOrEven(os.Args[i]))
		}

}

func oddOrEven(value string) string{
	num, _ := strconv.Atoi(value)
		if num % 2 == 0 {
			return "even"
		} else{
			return "odd"
		}
}

