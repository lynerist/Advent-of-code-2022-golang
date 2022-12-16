package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct{
	x,y int
}

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	line2000000 := make(map[int]bool)
	
	for sc.Scan(){
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)

		distanceFromLine := (sensorY -2000000)
		if distanceFromLine < 0 {
			distanceFromLine *= -1
		}

		for i:=0; i<= distanceFromBeacon - distanceFromLine; i++{
			line2000000[sensorX+i] = true
			line2000000[sensorX-i] = true
		}

		if beaconY == 2000000 {
			delete(line2000000, beaconX)
		}
	}
	fmt.Println(len(line2000000))

}

func manhattan(sensorX, sensorY, beaconX, beaconY int)int{
	x := sensorX - beaconX 
	y := sensorY - beaconY
	if x<0{
		x *= -1
	}
	if y<0{
		y *= -1
	}
	return x+y
}