package main

import(
	"fmt"
	"strconv"
)

func add(x,y int) int {
	return x + y
}

func subs(x,y int) int {
	return x - y
}

func div(x,y int) int {
	return x/y
}

func multi(x,y int) int {
	return x*y
}

func main() {
	fmt.Printf("Hello World!\n")
	
	fmt.Println("Operasi aritmatika:")
	fmt.Println("Penjumlahan: " + strconv.Itoa(add(10,20)))
	fmt.Println("Pengurangan: " + strconv.Itoa(subs(1,2)))
	fmt.Println("Pembagian: " + strconv.Itoa(div(2,3)))
	fmt.Println("Perkalian: " + strconv.Itoa(multi(2,3)))
}