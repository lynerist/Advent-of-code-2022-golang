package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct{
	x,y int
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	cave := make(map[point]rune)

	var maxY, maxX, minX int = 0,0,500

	//Generate cave
	for sc.Scan(){
		points := strings.Split(sc.Text(), " -> ")
		for i := range points[:len(points)-1]{
			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1]) 
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1]) 

			cave[point{toX, toY}] = '#'
			cave[point{fromX, fromY}] = '#'
			if toY>maxY{
				maxY = toY
			}
			if toX>maxX{
				maxX = toX
			}
			if toX<minX{
				minX = toX
			}

			for fromX != toX || fromY != toY{
				cave[point{fromX, fromY}] = '#'
				switch {
				case fromX < toX:
					fromX++
				case fromY < toY :
					fromY++
				case fromX > toX:
					fromX--
				case fromY > toY:
					fromY--
				}
			}
			if fromY>maxY{
				maxY = fromY
			}
			if fromX>maxX{
				maxX = fromX
			}
		}
	}
	for i := minX-500; i< maxX+500; i++{
		cave[point{i, maxY+2}]='#'
	}

	var sand int = 0
	for {
		newSand := point{500, 0}
		if cave[newSand]=='o'{
			break
		}
		for {
			cave[newSand]='"'
			if cave[point{newSand.x, newSand.y+1}] < '#'{
				newSand.y++
			}else if cave[point{newSand.x-1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x--
			}else if cave[point{newSand.x+1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x++
			}else{
				cave[newSand] = 'o'
				sand++
				break
			}
		} 
	}
	fmt.Println(sand)
	
	/* print the sand
	
	for y := 0; y<=maxY+5; y++{
		for x:= 400; x<=maxX+30; x++{
			if cave[point{x,y}] == 0{
				fmt.Print(" ")
			}else{
				fmt.Print(string(cave[point{x,y}]))
			}
		}
		fmt.Println()
	}*/
}