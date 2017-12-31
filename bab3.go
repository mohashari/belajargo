package main

import "fmt"

func main() {

	var s []int

	var i int
	for o := 0; o < 10; o++ {
		fmt.Println("masukan angka")
		fmt.Scanln(&i)
		s = append(s, i)
		printSlice(s)
	}

}

func printSlice(s []int) {
	fmt.Println("Result: ", len(s), cap(s), s)
}
