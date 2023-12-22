package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(Beam(str))
	// fmt.Println(BonusBeam(str))
}

func Beam(f string) (ans int){
	mirror := strings.Split(f, "\n")
	mirror = mirror[:len(mirror)-1]
	height := len(mirror)
	highestPoint := make([]int, len(mirror[0]))
	// go through the list. If we have a roller, push it to the highestPt
	for y, row := range mirror{
		for x, char := range row{
			// if it's a stationary rock the earliest you can stop is 1 below it
			switch char {
			case '#':
					highestPoint[x] = y+1
			case '.':
					continue
			case 'O':
				// roll the ball upwards and then add it in, block off its location.
				ans += (height-highestPoint[x])
				highestPoint[x] = highestPoint[x]+1
			default:
				fmt.Printf("issue with rune at x %d y %d \n", x, y)
			}
		}
	}
	return ans
}