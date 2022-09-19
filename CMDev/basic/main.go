package main

import (
	"fmt"
)

var num int = 10;

func add(a int, b int) int {
	return a + b + num;
}

func main() {
	var number int = 90;
	fmt.Println("Hello World.");
	fmt.Println("The number is", number);

	sum := add(1,2);
	fmt.Println(sum); 
	fmt.Println(num); 
}