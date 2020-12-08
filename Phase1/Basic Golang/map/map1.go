package main

import "fmt"

func main ()  {
	countrypoppulation := make (map [string] int )
	countrypoppulation = map [string] int {
		"India" : 123,
		"Usa" :33,

	}
	fmt.Println(countrypoppulation)
	// if you want to print particular one
	fmt.Println(countrypoppulation ["India"])
	// if you want to add value
	countrypoppulation ["Russia"] = 23
	fmt.Println(countrypoppulation)
	//if you want to delete
	delete(countrypoppulation , "Russia")
	fmt.Println(countrypoppulation)
}

