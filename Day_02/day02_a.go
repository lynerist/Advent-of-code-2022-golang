package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//A : X Rock 
	//B : Y Paper 
	//C : Z Scissors

	var score int
	scores := map[string]int{"B X":1, "C Y":2, "A Z":3, "A X":4,"B Y":5,"C Z":6, "C X":7, "A Y":8, "B Z":9}
	
	for sc.Scan(){
		score += scores[sc.Text()]	
	}
	fmt.Println(score)
}