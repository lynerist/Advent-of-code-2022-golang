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
	
	var forest [][]rune
	for sc.Scan(){
		row := []rune{}
		for _, tree := range sc.Text(){
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	var HighestScore int
	for i:=0; i<len(forest); i++{
		for j:=0; j<len(forest[0]);j++{
			score:=calcViewDown(forest,i,j,forest[i][j], true)*
			calcViewLeft(forest,i,j,forest[i][j], true)*
			calcViewRight(forest,i,j,forest[i][j], true)*
			calcViewTop(forest,i,j,forest[i][j], true)
			if score > HighestScore{
				HighestScore = score
			}
		}
	}
	fmt.Println(HighestScore)
}

func calcViewRight(forest [][]rune, i, j int, house rune, firstRec bool)int{
	if j == len(forest[0])-1 || !firstRec && forest[i][j]>=house{
		return 0
	}
	return 1 + calcViewRight(forest, i, j+1, house, false)
}

func calcViewLeft(forest [][]rune, i, j int, house rune, firstRec bool)int{
	if j == 0 || !firstRec && forest[i][j]>=house{
		return 0
	}
	return 1 + calcViewLeft(forest, i, j-1, house, false)
}

func calcViewDown(forest [][]rune, i, j int, house rune, firstRec bool)int{
	if i == len(forest)-1 || !firstRec && forest[i][j]>=house{
		return 0
	}
	return 1 + calcViewDown(forest, i+1, j, house, false) 
}

func calcViewTop(forest [][]rune, i, j int, house rune, firstRec bool)int{
	if i == 0 || !firstRec && forest[i][j]>=house{
		return 0
	}
	return 1 + calcViewTop(forest, i-1, j, house, false)
}