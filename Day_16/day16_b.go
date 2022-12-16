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
	elephantPosition string
	opened map[string]bool
	visitedWithReleasedPerMinute map[string]int
	visitedWithReleasedPerMinuteElephant map[string]int

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

	possibleStates := []state{state{0,0,"AA","AA", zeroRates, make(map[string]int),make(map[string]int)}}
	var finalPressures []int

	for minute := 0; minute<26; minute++{
		fmt.Printf("%d %8d",minute, len(possibleStates))
		var newPossibleStates []state
		for _, current := range possibleStates {

			if len(current.opened) == len(nodes){
				finalPressures = append(finalPressures, current.releasedPressure + current.releasedPerMinute* (26-minute))
				continue
			}

			if minute>6{
			//	fmt.Println(current)
			}

			//Parameters to tune to reach the best solution in acceptable time

			
			if minute > 7 && current.releasedPerMinute<30{
				continue
			}
			if minute > 9 && current.releasedPerMinute<60{
				continue
			}

			if minute > 10 && (current.releasedPressure<200){
				continue
			}		
			
			if minute > 13 && (current.releasedPressure<450 || current.releasedPerMinute<70){
				continue
			}

			if minute > 14 && (current.releasedPressure<550 || current.releasedPerMinute<75){
				continue
			}

			
			if minute > 15 && (current.releasedPressure<700 || current.releasedPerMinute<85){
				continue
			}
			
			
			if minute > 16 && (current.releasedPressure<800 || current.releasedPerMinute<100) {
				continue
			}
			
			if minute > 17 && (current.releasedPressure<900 ||current.releasedPerMinute<105){
				continue
			}
			
			if minute > 18 && (current.releasedPressure<1100 || current.releasedPerMinute<110){
				continue
			}
			
			if minute > 19 && (current.releasedPressure<1250 || current.releasedPerMinute<115){
				continue
			}

			if minute > 20 && (current.releasedPressure<1300 || current.releasedPerMinute<120){
				continue
			}
			
			if minute > 21 &&(current.releasedPressure<1400 || current.releasedPerMinute<125){
				continue
			}
			
			if minute > 23 &&(current.releasedPressure<1700 || current.releasedPerMinute<135){
				continue
			}
			if minute > 24 &&(current.releasedPressure<1900 || current.releasedPerMinute<140){
				continue
			}
			


			current.visitedWithReleasedPerMinute[current.position] = current.releasedPerMinute
			current.visitedWithReleasedPerMinuteElephant[current.elephantPosition] = current.releasedPerMinute
			
			//Combinations of moves mine and of the elephant

			if !current.opened[current.position] {
				var atLeastOne bool
				for _, tunnelForElephant := range nodes[current.elephantPosition].tunnels{
					if releasedPerMinute, ok := current.visitedWithReleasedPerMinuteElephant[tunnelForElephant]; ok && releasedPerMinute == current.releasedPerMinute{
						continue
					}
					newOpened := copyMapBool(current.opened)
					newOpened[current.position] = true
					newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																		current.releasedPerMinute+ nodes[current.position].flowRate,
																		current.position, tunnelForElephant ,newOpened, copyMapInt(current.visitedWithReleasedPerMinute),
																		copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
					atLeastOne = true
				}

				if !current.opened[current.elephantPosition] && current.position!=current.elephantPosition{
					newOpened := copyMapBool(current.opened)
					newOpened[current.position] = true
					newOpened[current.elephantPosition] = true
					newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																		current.releasedPerMinute+ nodes[current.position].flowRate + nodes[current.elephantPosition].flowRate,
																		current.position, current.elephantPosition ,newOpened, copyMapInt(current.visitedWithReleasedPerMinute),
																		copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
				} else if !atLeastOne{
					newOpened := copyMapBool(current.opened)
					newOpened[current.position] = true
					newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																		current.releasedPerMinute+ nodes[current.position].flowRate,
																		current.position, current.elephantPosition ,newOpened, copyMapInt(current.visitedWithReleasedPerMinute),
																		copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
				}
			}
			var atLeastOne bool
			for _, tunnel := range nodes[current.position].tunnels{
				if releasedPerMinute, ok := current.visitedWithReleasedPerMinute[tunnel]; ok && releasedPerMinute == current.releasedPerMinute{
					continue
				}
				for _, tunnelForElephant := range nodes[current.elephantPosition].tunnels{
					if releasedPerMinute, ok := current.visitedWithReleasedPerMinuteElephant[tunnelForElephant]; ok && releasedPerMinute == current.releasedPerMinute{
						continue
					}
					//fmt.Println(tunnel, current.releasedPerMinute, current.opened)
					newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																		current.releasedPerMinute, tunnel, tunnelForElephant, copyMapBool(current.opened),
																		copyMapInt(current.visitedWithReleasedPerMinute),copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
				}
				if !current.opened[current.elephantPosition]{
					newOpened := copyMapBool(current.opened)
					newOpened[current.elephantPosition] = true
					newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																		current.releasedPerMinute+ nodes[current.elephantPosition].flowRate,
																		tunnel, current.elephantPosition ,newOpened, copyMapInt(current.visitedWithReleasedPerMinute),
																		copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
				}
				atLeastOne = true
			}
			if !atLeastOne && !current.opened[current.elephantPosition]{
				newOpened := copyMapBool(current.opened)
				newOpened[current.elephantPosition] = true
				newPossibleStates = append(newPossibleStates, state{current.releasedPressure + current.releasedPerMinute,
																	current.releasedPerMinute+ nodes[current.elephantPosition].flowRate,
																	current.position, current.elephantPosition ,newOpened, copyMapInt(current.visitedWithReleasedPerMinute),
																	copyMapInt(current.visitedWithReleasedPerMinuteElephant)})
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

	for _, releasedPressure := range finalPressures{
		if releasedPressure > maxPressure{
			maxPressure = releasedPressure
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