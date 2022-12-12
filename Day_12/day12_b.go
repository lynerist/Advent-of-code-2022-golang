package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct{
	x,y int
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	heightmap := make([][]rune, 0)
	var end point
	var starts []point

	for sc.Scan(){
		var line []rune
		for i, elevation := range sc.Text(){
			if elevation == 'S' || elevation == 'a'{
				elevation = 'a'
				starts = append(starts, point{i,len(heightmap)})
			}
			if elevation == 'E'{
				end = point{i, len(heightmap)}
				elevation = 'z'
			}
			line = append(line, elevation)
		}
		heightmap = append(heightmap, line)	
	}
	var shortestPath int
	for _, start := range starts{
		visited := make(map[point]bool)
		toVisit := []point{start}
		distanceFromStart := map[point]int{start:0}

		for {	
			if len(toVisit) == 0 {
				break
			}
			currentPoint := toVisit[0]
			visited[currentPoint] = true
			toVisit = toVisit[1:]

			if currentPoint == end{
				if distanceFromStart[end]<shortestPath || shortestPath == 0{
					shortestPath = distanceFromStart[end]
				}
				break
			}
			
			for _, near := range [][]int{{1,0},{0,-1},{-1,0},{0,1}}{
				j, i := near[1], near[0]
				nextPoint := point{currentPoint.x+j, currentPoint.y+i} 
				if !visited[nextPoint] && nextPoint.x>=0 && nextPoint.y>=0 && 
				nextPoint.x<len(heightmap[0]) && nextPoint.y<len(heightmap) &&
				(heightmap[nextPoint.y][nextPoint.x]-heightmap[currentPoint.y][currentPoint.x]<=1){
					if (distanceFromStart[nextPoint] == 0){
						toVisit = append(toVisit, nextPoint)
						distanceFromStart[nextPoint]=distanceFromStart[currentPoint]+1
					}
					if (distanceFromStart[nextPoint]>=distanceFromStart[currentPoint]+1){
						distanceFromStart[nextPoint]=distanceFromStart[currentPoint]+1
					}
				}	
			}
			sort.Slice(toVisit, func(i,j int)bool{
				return distanceFromStart[toVisit[i]] < distanceFromStart[toVisit[j]]
			})
		}
	}
	fmt.Println(shortestPath)
}