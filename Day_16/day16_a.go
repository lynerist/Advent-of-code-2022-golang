package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct{
	flowRate int
	tunnels []string
}

type state struct{
	releasedPressure int
	releasedPerMinute int
	position string
	opened map[string]bool
	visitedWithReleasedPerMinute map[string]int
}

func (s state) String()string{
	return fmt.Sprintf("%d per min,  %d released %d opened", s.releasedPerMinute, s.releasedPressure, len(s.opened))
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	nodes := make(map[string]node)
	zeroRates := make(map[string]bool)

	for sc.Scan(){
		input := strings.Fields(sc.Text())
		var newNode node
		fmt.Sscanf(input[4], "rate=%d;", &newNode.flowRate)

		for _, tunnel := range input[9:]{
			newNode.tunnels = append(newNode.tunnels, tunnel[:2])
		}
		nodes[input[1]] = newNode

		if newNode.flowRate == 0{
			zeroRates[input[1]] = true
		}
	}

	possibleStates := []state{state{0,0,"AA",zeroRates, make(map[string]int)}}

	for minute := 0; minute<30; minute++{
		var newPossibleStates []state
		for _, current := range possibleStates {

			//Parameters to tune to reach the best solution in acceptable time
			if minute>12 && (current.releasedPerMinute<40){
				continue
			}
			if minute > 24 && current.releasedPressure<1000{
				continue				 
			}
						
			current.visitedWithReleasedPerMinute[current.position] = current.releasedPerMinute

			if !current.opened[current.position] {
				newOpened := copyMapBool(current.opened)
				newOpened[current.position] = true
				newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																	current.releasedPerMinute+ nodes[current.position].flowRate,
																	current.position, newOpened, copyMapInt(current.visitedWithReleasedPerMinute)})
			}
			for _, tunnel := range nodes[current.position].tunnels{

				if releasedPerMinute, ok := current.visitedWithReleasedPerMinute[tunnel]; ok && releasedPerMinute == current.releasedPerMinute{
					continue
				}
				newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																	current.releasedPerMinute, tunnel, copyMapBool(current.opened),
																	copyMapInt(current.visitedWithReleasedPerMinute)})
			}
		}
		possibleStates = newPossibleStates
	}

	var maxPressure int
	
	for _, possibleState := range possibleStates{
		if possibleState.releasedPressure > maxPressure{
			maxPressure = possibleState.releasedPressure
		}
	}

	fmt.Println(maxPressure)
}

func copyMapBool(mapToCopy map[string]bool)map[string]bool{
	newMap := make(map[string]bool)
	for key, value := range mapToCopy{
		newMap[key] = value
	}
	return newMap
}

func copyMapInt(mapToCopy map[string]int)map[string]int{
	newMap := make(map[string]int)
	for key, value := range mapToCopy{
		newMap[key] = value
	}
	return newMap
}