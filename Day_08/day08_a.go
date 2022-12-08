package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x,y int
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	var forest [][]rune
	for sc.Scan(){
		row := []rune{}
		for _, tree := range sc.Text(){
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	maxLeft := make([]rune, len(forest))
	maxRigth := make([]rune, len(forest))

	//Horizontal view
	isVisible := make(map[point]bool)
	for i:=0; i<len(forest); i++{
		for j:=0; j<len(forest[0]);j++{
			if j==0{
				maxLeft[i] = forest[i][j]
				maxRigth[i]= forest[i][len(forest[0])-1]
				isVisible[point{i,j}]=true
				isVisible[point{i,len(forest[0])-1}]=true
				continue
			}
			if forest[i][j]>maxLeft[i]{
				isVisible[point{i,j}]=true
				maxLeft[i]=forest[i][j]
			}
			if forest[i][len(forest[0])-1-j]>maxRigth[i]{
				isVisible[point{i,len(forest[0])-1-j}] = true
				maxRigth[i] = forest[i][len(forest[0])-1-j]
			}
		}
	}
	
	//Vertical view

	maxTop := make([]rune, len(forest))
	maxDown := make([]rune, len(forest))

	for i:=0; i<len(forest); i++{
		for j:=0; j<len(forest[0]);j++{
			if i==0{
				maxTop[j] = forest[i][j]
				maxDown[j]= forest[len(forest)-1][j]
				isVisible[point{i,j}]=true
				isVisible[point{len(forest)-1,j}]=true
				continue
			}
			if forest[i][j]>maxTop[j]{
				isVisible[point{i,j}]=true
				maxTop[j]=forest[i][j]
			}
			if forest[len(forest)-1-i][j]>maxDown[j]{
				isVisible[point{len(forest)-1-i,j}] = true
				maxDown[j] = forest[len(forest)-1-i][j]
			}
		}
	}
	fmt.Println(len(isVisible))
}
