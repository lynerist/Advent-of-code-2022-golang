package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct{
	items []int
	operation func(int)int
	testAndThrow func(int)int
} 

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var monkeys []monkey

	//Parsing input
	for sc.Scan(){
		sc.Scan()
		var newMonkey monkey
		for _, item := range strings.Split(sc.Text()[len("  Starting items: "):], ", "){
			worryLevel,_ := strconv.Atoi(item)
			newMonkey.items = append(newMonkey.items, worryLevel)
		}

		sc.Scan()
		var operator rune
		var value int = 0
		if sc.Text() == "  Operation: new = old * old" || sc.Text()=="  Operation: new = old + old"{
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c old", &operator)
		}else{
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c %d", &operator, &value)
		}
		newMonkey.operation = createOperation(operator, value)

		sc.Scan()
		var testingValue int
		fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &testingValue)

		sc.Scan()
		var toThrowIfTrue int
		fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &toThrowIfTrue)	
		sc.Scan()
		var toThrowIfFalse int
		fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &toThrowIfFalse)
		newMonkey.testAndThrow = createTestAndThrow(testingValue, toThrowIfTrue, toThrowIfFalse)

		sc.Scan()
		monkeys = append(monkeys, newMonkey)
	}

	counts := make([]int, len(monkeys))

	for turn := 0; turn < 20; turn++{
		for monkeyId, currentMonkey := range monkeys{
			for _, item := range currentMonkey.items{
				newValue := currentMonkey.operation(item)
				monkeys[currentMonkey.testAndThrow(newValue)].items = append(monkeys[currentMonkey.testAndThrow(newValue)].items, newValue)
			}
			counts[monkeyId] += len(monkeys[monkeyId].items)
			monkeys[monkeyId].items = []int{}
		}
	}

	var highestCount, secondHighest int
	for _, count := range counts{
		if count > secondHighest{
			secondHighest = count
		}
		if secondHighest>highestCount{
			highestCount, secondHighest = secondHighest, highestCount
		}
	}
	fmt.Println(highestCount*secondHighest)
}

func createOperation(operator rune, value int)(func(int)int){
	operation := func (n int)int{
		var valueToUse int = value
		if valueToUse == 0{
			valueToUse = n
		}
		if operator == '+'{
			return (n+valueToUse)/3
		}
		return (n*valueToUse)/3
	}
	return operation
}

func createTestAndThrow(testingValue, toThrowIfTrue, toThrowIfFalse int)(func(int)int){
	testAndThrow := func(n int)int{
		if n%testingValue==0{
			return toThrowIfTrue
		}
		return toThrowIfFalse
	}
	return testAndThrow
}