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
	
	const limit int = 4000000
	
	toTry := make(map[point]bool)
	nearestBeacon := make(map[point]int)
	
	for sc.Scan(){
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)
		nearestBeacon[point{sensorX,sensorY}] = distanceFromBeacon
		distanceFromBeacon++
		for i:=0; i<distanceFromBeacon; i++{
			if (sensorX+i>0 && sensorX+i<limit){
				if (sensorY-distanceFromBeacon+1+i> 0 &&  sensorY-distanceFromBeacon+1+i<limit){
					toTry[point{sensorX+i, sensorY-distanceFromBeacon+1+i}] = true
				}
				if (sensorY+distanceFromBeacon-1-i> 0 && sensorY+distanceFromBeacon-1-i<limit){
					toTry[point{sensorX+i, sensorY+distanceFromBeacon-i}] = true
				}
			}
			if (sensorX-i>0 && sensorX-i<limit){
				if ( sensorY-distanceFromBeacon+1+i> 0 &&  sensorY-distanceFromBeacon+1+i<limit){
					toTry[point{sensorX-i,  sensorY-distanceFromBeacon+1+i}] = true
				}
				if (sensorY+distanceFromBeacon-1-i> 0 && sensorY+distanceFromBeacon-1-i<limit){
					toTry[point{sensorX-i, sensorY+distanceFromBeacon-1-i}] = true
				}
			}
		}		
	}
	
	for beacon := range toTry {
		newBeacon := true
		for sensor, nearestBeaconDistance := range nearestBeacon{
			if manhattan(sensor.x, sensor.y, beacon.x, beacon.y) <= nearestBeaconDistance{
				newBeacon = false
				break
			}
		}
		if newBeacon {
			fmt.Println(beacon.x*limit+beacon.y)
			break
		}
	}
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