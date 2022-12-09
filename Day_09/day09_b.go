package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x,y int
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	visitedByTail := make(map[point]bool)
	knots := make([]point, 10)

	visitedByTail[knots[9]]=true
	
	for sc.Scan(){
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])
		//I calculate moves one by one
		for moves > 0{
			switch direction{
			case 'U':
				knots[0].y++
			case 'R':
				knots[0].x++
			case 'D':
				knots[0].y--
			case 'L':
				knots[0].x--
			}
			//Each knot is the head of the successive knot
			for i := range knots[:len(knots)-1]{
				knots[i+1] = adjustTail(knots[i+1], knots[i])
			}
			moves--
			visitedByTail[knots[9]]=true
		}
	}

	fmt.Println(len(visitedByTail))
}

func adjustTail(tail point, head point)(newTail point){
	newTail = tail
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,1},point{-1,2}, point{0,2}, point{1,2}, point{2,1}, point{2,2}, point{-2,2}:
		newTail.y++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{1,2},point{2,1}, point{2,0}, point{2,-1}, point{1,-2}, point{2,2}, point{2,-2}:
		newTail.x++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{2,-1},point{1,-2}, point{0,-2}, point{-1,-2}, point{-2,-1}, point{2,-2}:
		newTail.y--
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{-1,-2},point{-2,-1}, point{-2,-0}, point{-2,1}, point{-1,2}, point{-2,2}:
		newTail.x--
	}
	return
}
