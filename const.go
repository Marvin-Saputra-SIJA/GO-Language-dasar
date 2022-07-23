//const in go is simmiliar like varibale but the value cannot be change after it declared

package main

import "fmt"

func main () {
	
	const (
		KELAS = "11 SIJA"
		MAPEL = "Software Computer"		
	)
	var nama = "Marvin Saputra"
	
	fmt.Println(nama,KELAS,MAPEL)
}