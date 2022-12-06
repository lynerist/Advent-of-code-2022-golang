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
	sc.Scan()
	
	const differentCharactersNeeded = 4
	for i := range sc.Text() {
		charactersSet := make(map[byte]bool)
		for j:=0; j<differentCharactersNeeded; j++{
			charactersSet[sc.Text()[i+j]]=true
		}
		if len(charactersSet) == differentCharactersNeeded{
			fmt.Println(i+differentCharactersNeeded)
			break
		}
	}
}
