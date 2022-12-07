package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name string
	size int
	isFile bool
	sons map[string]*node
	father *node
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	var currentDirectory *node

	dirs := []*node{}

	for sc.Scan(){
		line := strings.Fields(sc.Text())
		if len(line)>2{
			if line[2]==".."{
				currentDirectory = currentDirectory.father
			}else if line[2]=="/"{
				currentDirectory = &node{"/", 0, false, make(map[string]*node), nil}
				dirs = append(dirs, currentDirectory)
			}else{
				currentDirectory = currentDirectory.sons[line[2]]
			}
		}else if line[0] == "dir"{
			currentDirectory.sons[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
			dirs = append(dirs, currentDirectory.sons[line[1]])
		}else if line[0] != "$"{
			size, _ := strconv.Atoi(line[0])
			currentDirectory.sons[line[1]] = &node{line[1], size, true, nil, currentDirectory}
		}
	}

	toFree:= 30000000 - (70000000 - calcSize(*dirs[0]))
	var smallestEnaugthSize int = calcSize(*dirs[0])

	for _, dir := range dirs{
		size := calcSize(*dir)
		if size > toFree && size-toFree < smallestEnaugthSize-toFree{
			smallestEnaugthSize = size
		}
	}

	fmt.Println(smallestEnaugthSize)
}


func calcSize(root node)(size int){
	if root.isFile {
		return root.size
	} 
	for _, d := range root.sons{
		size += calcSize(*d)
	}
	return
}
