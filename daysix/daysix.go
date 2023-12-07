package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(Boat(str))
	fmt.Println(BonusBoat(str))
}

func Boat(f string) (ans int){
	// make two slices of all the numbers
	timeStr, distStr, found := strings.Cut(f, "\n")
	if found == false{
		fmt.Println("incorrect input!")
		return 0
	}
	var time []int
	var dist []int
	for _, num := range strings.Fields(timeStr){
		if i, err := strconv.Atoi(num); err == nil{
			time = append(time, i)
		}
	}
	for _, num := range strings.Fields(distStr){
		if i, err := strconv.Atoi(num); err == nil{
			dist = append(dist, i)
		}
	}
	ans = 1
	for key := range time{
		t := time[key]
		d := dist[key]
		halft := t/2
		for i:= 1; i<=halft; i++{
			if i*(t-i) >d{
				//if you beat the record, you need to add the number of winning speeds
				ans *= (halft-i)*2 + (t%2) +1
				break
			}
		}
	}
	// start at 1, int-1, and go until the answer beats the time
	// once the number beats the set time multiply by 2 and add to answer
	return ans
}

func BonusBoat(f string) (ans int){
	// make two slices of all the numbers
	timeStr, distStr, found := strings.Cut(f, "\n")
	if found == false{
		fmt.Println("incorrect input!")
		return 0
	}
	var time string
	var dist string
	for _, num := range strings.Fields(timeStr){
		if _, err := strconv.Atoi(num); err == nil{
			time += num
		}
	}
	for _, num := range strings.Fields(distStr){
		if _, err := strconv.Atoi(num); err == nil{
			dist += num
		}
	}
	ans = 1
	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(dist)
	halft := t/2
	for i:= 1; i<=halft; i++{
		if i*(t-i) >d{
			//if you beat the record, you need to add the number of winning speeds
			ans *= (halft-i)*2 + (t%2) +1
			break
		}
	}
	// start at 1, int-1, and go until the answer beats the time
	// once the number beats the set time multiply by 2 and add to answer
	return ans
}