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

	maxCalory := 0
	currentCalory := 0

	for sc.Scan(){
		calories, err := strconv.Atoi(sc.Text())
		currentCalory += calories

		if err != nil{
			if currentCalory>maxCalory{
				maxCalory = currentCalory
			}
			currentCalory = 0
		}
	}
	fmt.Println(maxCalory)
}