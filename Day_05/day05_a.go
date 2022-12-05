package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	elements []rune
}

func (s *stack) push(r rune){
	s.elements = append(s.elements, r)
}

func (s *stack) pop()(r rune){
	r = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *stack) addToBottom(r rune){
	s.elements = append([]rune{r}, s.elements...)
}

func (s stack) String()string{
	var str string
	for _, r := range s.elements{
		str += string(r)+" "
	}
	return str
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//create slice of stacks
	stacks := make([]stack, 9)

	//Parsing the input
	sc.Scan()
	for (sc.Text()!=" 1   2   3   4   5   6   7   8   9 "){
		for i, r := range sc.Text(){
			if r != ' ' && r!='[' && r!=']'{
				stacks[i/4].addToBottom(r)
			}
		}
		sc.Scan()
	}
	//Read empty line
	sc.Scan()

	for sc.Scan(){
		var toMove, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

		//Move elements one by one
		for move :=0; move < toMove; move++{
			stacks[to-1].push(stacks[from-1].pop())
		}
	}

	for _, s := range stacks{
		fmt.Print(string(s.pop()))
	}

}