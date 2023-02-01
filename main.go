package main

import (
	"fmt"
	//"log"
)

func area(length float64, width float64) float64{
	result := length * width
	return result 
}

func main() {
	lenght := 5.5
	width := -10
	//call the area function
	result := area(lenght, float64(width))
	fmt.Println(result)
}

