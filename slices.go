//slices is simmiliar to array but you dont have to define lengt of values. Its more flexible and powerfull than array

package main

import "fmt"

func main () {

	var sija1 = []string{"Aris","Mila"}
	var sija2 = []string{"Marvin","Ilham","Farhan"}

	sija2[2] = "Fiona" //using to modify slices
	fmt .Println(sija2[2]) //print values of slices
	fmt.Println(len(sija2[2]))

	var sija3 = append(sija1, sija2...) //using append function to concate 2 slices into 1 slices 
	fmt.Println(sija3)
}
