package main

import (
	"fmt"
	"log"
	"errors"
)

func area(length float64, width float64) (float64, error){
	//check if lenght is greater than 0 
	if length < 0 {
		err := errors.New("you cannot have a negative lenght")
		return -1, err 
	}
	if width < 0 {
		err := errors.New("you cannot have a negative width")
		return -1, err 
	}

	result := length * width
	return result, nil
}

func main() {
	lenght := 5.5
	width := 10
	//call the area function
	result, err := area(lenght, float64(width))
	//fmt.Println(result)
	if err != nil{
		log.Fatal(err)
	}
	fmt. Println(result)
}

