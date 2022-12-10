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
	head := point{0,0}
	tail := point{0,0}
	visitedByTail[tail]=true
	for sc.Scan(){
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])

		//I calculate moves one by one
		for moves > 0{
			switch direction{
			case 'U':
				head.y++
			case 'R':
				head.x++
			case 'D':
				head.y--
			case 'L':
				head.x--
			}
			moves--
			tail = adjustTail(tail, head)
			visitedByTail[tail]=true
		}
	}

	fmt.Println(len(visitedByTail))
}

func adjustTail(tail point, head point)(newTail point){
	newTail = tail
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,1},point{-1,2}, point{0,2}, point{1,2}, point{2,1}:
		newTail.y++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{1,2},point{2,1}, point{2,0}, point{2,-1}, point{1,-2}:
		newTail.x++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{2,-1},point{1,-2}, point{0,-2}, point{-1,-2}, point{-2,-1}:
		newTail.y--
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-1,-2},point{-2,-1}, point{-2,-0}, point{-2,1}, point{-1,2}:
		newTail.x--
	}
	return
}
