package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	file,_ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	maxCalory1 := 0
	maxCalory2 := 0
	maxCalory3 := 0

	currentCalory := 0

	for sc.Scan(){
		calories, err := strconv.Atoi(sc.Text())
		currentCalory += calories

		if err != nil{
			if currentCalory>maxCalory3{
				maxCalory3 = currentCalory
			}
			if maxCalory3>maxCalory2{
				maxCalory3,maxCalory2 = maxCalory2, maxCalory3
			}
			if maxCalory2>maxCalory1{
				maxCalory2,maxCalory1 = maxCalory1, maxCalory2
			}
			currentCalory = 0
		}
	}
	fmt.Println(maxCalory1+maxCalory2+maxCalory3)
}