//operator in golang 
//theres a few operator in golang 

//arithmetic
//assignment
//comparison 
//logic
//bitwise 

//arithmetic
package main

import "fmt"

func main () {
	fmt.Println("==============arithmetic===========")
	var x = 10;
	x += 5; // add values using operator 
	fmt.Println(x)
	fmt.Println("==============comparison===========")
    x = 11;
	var z = 12;
	fmt.Println(x == z);
	fmt.Println("==============logic===========")
	x = 10;
	var y = 20;

	fmt.Println(x < y || x > y);
}
