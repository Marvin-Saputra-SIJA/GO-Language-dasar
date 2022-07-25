// condition in go using if else, else if and switch statement 
package main 

import "fmt"

func main () {


	var (
		// name = "marvin";
		grade = 81
	)
	 
	// if grade == 75  {
	// 	fmt.Printf(" You can apply for re-exam your score is %d ", grade)
	// 	} else if grade <= 75  {
	// 		fmt.Printf("Im sorry %s, you didnt pass the exam, your score is %d ", name, grade)
	// } else {
	// 	fmt.Printf("Congratulations %s, you pass the exam, your score is %d ", name, grade)
	// }

	switch grade {
		case 75:
			fmt.Printf("re exam")
		case 80,81,82:
			fmt.Printf("pass the exam")
		case 70:
			fmt.Printf("didnt pass the exam")
	}
}


