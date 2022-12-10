package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	registerX, cycleNumber := 1, 0
	var finalValue int

	for sc.Scan(){
		operations := strings.Fields(sc.Text())
		incrementAndControl(&cycleNumber, &registerX, &finalValue)
		if operations[0]=="addx"{
			value, _ := strconv.Atoi(operations[1])
			incrementAndControl(&cycleNumber, &registerX, &finalValue)
			registerX += value
		}
	}
	fmt.Println(finalValue)
}

func incrementAndControl(cycleNumber, registerX, finalValue *int){
	*cycleNumber++
	if (*cycleNumber-20)%40==0 && *cycleNumber<=220{
		*finalValue += *registerX * *cycleNumber
	}
}