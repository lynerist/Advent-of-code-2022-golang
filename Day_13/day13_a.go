package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tree struct{
	valueLeaf int
	elements []*tree
	father *tree
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	index := 1
	var indexSum int
	for sc.Scan(){
		package1 := readTree(sc.Text())				
		sc.Scan()
		package2 := readTree(sc.Text())		
		
		if areOrdered(package1, package2) == 1 {
			indexSum += index
		}

		index++
		sc.Scan()
	}
	fmt.Println(indexSum)
}

func readTree(input string)(tree){
	root := tree{-1, []*tree{}, nil}
	temp := &root
	
	var currentNumber string
	for _, r := range input{
		switch r {
		case '[':
			newTree := tree{-1, []*tree{}, temp}
			temp.elements=append(temp.elements, &newTree)
			temp = &newTree
		case ']':
			if len(currentNumber)>0{
				number, _ := strconv.Atoi(currentNumber)
				temp.valueLeaf = number
				currentNumber = ""
			}
			temp = temp.father
		case ',':
			if len(currentNumber)>0{
				number, _ := strconv.Atoi(currentNumber)
				temp.valueLeaf = number
				currentNumber =""
			}
			temp = temp.father
			newTree := tree{-1, []*tree{}, temp}
			temp.elements=append(temp.elements, &newTree)
			temp = &newTree
		default:
			currentNumber += string(r)
		}
	}
	return root
}

func areOrdered(first, second tree)int{
	switch{
	case len(first.elements) == 0 && len(second.elements) == 0:
		if first.valueLeaf > second.valueLeaf{
			return -1
		} else if first.valueLeaf == second.valueLeaf {
			return 0
		}
		return 1
		
	case first.valueLeaf >= 0:
		return areOrdered(tree{-1, []*tree{&first}, nil}, second)

	case second.valueLeaf >= 0:
		return areOrdered(first, tree{-1, []*tree{&second}, nil})
	default:
		var i int
		for i=0; i<len(first.elements) && i<len(second.elements); i++{
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0{
				return ordered
			}
		}
		if i < len(first.elements){
			return -1
		}else if i < len(second.elements){
			return 1
		}
	}
	return 0
}