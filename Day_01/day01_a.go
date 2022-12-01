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

	//Search for the maximum sum of calories
	maxCalories := 0
	currentElfCalories := 0

	for sc.Scan(){
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		//If error is different from nil then I found an empty line
		if err != nil{
			if currentElfCalories>maxCalories{
				maxCalories = currentElfCalories
			}
			//I start with a new elf
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories)
}