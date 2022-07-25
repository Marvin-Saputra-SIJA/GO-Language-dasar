//looping in go using for 

package main 

import "fmt"

func main () {
	var name = "marvin"
	for x := 1; x <= 10; x++ {
		if x % 2 == 1 {
			continue
		}
		fmt.Printf("%d hello %s\n",x, name )
	}
}