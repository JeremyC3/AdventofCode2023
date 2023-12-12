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
	fmt.Println(Pipe(str))
	// fmt.Println(BonusPipe(str))
}

func Pipe(f string)(ans int){
	// step 1 is to turn this into a 2d grid
	rows := strings.Split(f, "\n")
	// the last value is an empty string, remove it
	rows = rows[:len(rows)-1]
	height := len(rows)
	width := len(rows[0])
	sketch := make([][]rune, height)
	for i:= range sketch{
		sketch[i] = make([]rune, width)
	}
	type coord struct{
		X int
		Y int
	}
	// record the location of the animal too
	animalPos := coord{0,0}
	for y, row := range rows{
		for x, point := range row{
			sketch[y][x] = point
			if point == 'S'{
				animalPos = coord{x, y}
			}
		}
	}
	cardinals := [4]rune{'N', 'E','S','W'}
	CardinalCheck:
	for _, prevDir := range cardinals{
		ans = 0
		currPos := animalPos
		fmt.Printf("starting in direction %s \n",string(prevDir) )
		switch prevDir{
		case 'N':
			currPos.Y -= 1
		case 'E':
			currPos.X += 1
		case 'S':
			currPos.Y += 1
		case 'W':
			currPos.X -= 1
		}
		for(currPos.X != animalPos.X || currPos.Y != animalPos.Y){
			fmt.Printf("we are at X %d, Y %d\n", currPos.X, currPos.Y)
			// if we're outside of our range bail back to the start
			if currPos.X >= width || currPos.Y >= height || currPos.X < 0 || currPos.Y < 0{
				continue CardinalCheck
			}
			// get the current rune from the sketch
			ans++
			currRune := sketch[currPos.Y][currPos.X]
			switch prevDir{
			case 'N':
				switch currRune{
				case '|':
					prevDir = 'N'
					currPos.Y -= 1
				case 'F':
					prevDir = 'E'
					currPos.X += 1
				case '7':
					prevDir = 'W'
					currPos.X -=1
				default:
					continue CardinalCheck
				}
			case 'E':
				switch currRune{
				case 'J':
					prevDir = 'N'
					currPos.Y -= 1
				case '-':
					prevDir = 'E'
					currPos.X += 1
				case '7':
					prevDir = 'S'
					currPos.Y +=1
				default:
					continue CardinalCheck
				}
			case 'S':
				switch currRune{
				case 'J':
					prevDir = 'W'
					currPos.X -= 1
				case 'L':
					prevDir = 'E'
					currPos.X += 1
				case '|':
					prevDir = 'S'
					currPos.Y +=1
				default:
					continue CardinalCheck
				}
			case 'W':
				switch currRune{
				case 'L':
					prevDir = 'N'
					currPos.Y -= 1
				case '-':
					prevDir = 'W'
					currPos.X -= 1
				case 'F':
					prevDir = 'S'
					currPos.Y +=1
				default:
					continue CardinalCheck
				}
			}
		}
		if (currPos.X== animalPos.X && currPos.Y== animalPos.Y){
			fmt.Println("finished")
			fmt.Println(ans)
			return ans/2 + ans%2
		}
	}
	// for each of the cardinal directions, try the recursive function?
	// keep track of the last location + update currPipe
	return -1
}