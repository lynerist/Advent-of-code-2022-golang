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
		items := make(map[rune]bool)
		// Create a set with all the elements of the first compartment
		for _, itemLeftPart := range sc.Text()[:len(sc.Text())/2]{
			items[itemLeftPart] = true
		}

		// Range all the items of the second compartment
		for _, itemRightPart := range sc.Text()[len(sc.Text())/2:]{
			// If an item is in the first set it's in both compartments
			if items[itemRightPart]{
				sumOfPriorities += int(unicode.ToLower(itemRightPart)-96)
				if unicode.IsUpper(itemRightPart){
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}