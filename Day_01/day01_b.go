package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	maxCalories1 := 0
	maxCalories2 := 0
	maxCalories3 := 0

	currentElfCalories := 0

	for sc.Scan(){
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		//If error is different from nil then I found an empty line
		if err != nil{
			if currentElfCalories>maxCalories3{
				maxCalories3 = currentElfCalories
			}
			if maxCalories3>maxCalories2{
				maxCalories3,maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2>maxCalories1{
				maxCalories2,maxCalories1 = maxCalories1, maxCalories2
			}
			//I start with a new elf
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories1+maxCalories2+maxCalories3)
}