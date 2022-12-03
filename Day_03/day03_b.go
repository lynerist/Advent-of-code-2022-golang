package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sumOfPriorities int

	for sc.Scan() {
		// Create three sets with the elements of each elf
		itemsFirst := createSetOfItems(sc.Text())
		sc.Scan()
		itemsSecond := createSetOfItems(sc.Text())
		sc.Scan()
		itemsThird := createSetOfItems(sc.Text())

		// For each item of the first (or second or third, it's the same) elf we check if the other two elves have that item too
		for itemFirstElf := range itemsFirst{
			if itemsSecond[itemFirstElf] && itemsThird[itemFirstElf]{
				sumOfPriorities += int(unicode.ToLower(itemFirstElf)-96)
				if unicode.IsUpper(itemFirstElf){
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

func createSetOfItems(items string)(set map[rune]bool){
	set = make(map[rune]bool)
	for _, item := range items{
		set[item] = true
	}
	return
}