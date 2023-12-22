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
	fmt.Println(Lava(str))
	// fmt.Println(BonusLava(str))
}
// take in a pointer to the arr and also new coord + directions?

func Lava(f string)(ans int){
	// create a structure to store rune + direction
	
	type node struct{
		char rune
		pastDir string
	}
	graph := [][]node{}
	// first, turn the entire string into a 2d slice of slices
	for _, row := range strings.Split(f, "\n"){
		if row == ""{
			continue
		}
		nextCol := []node{}
		for _, char := range row{
			nextCol = append(nextCol, node{char, ""})
		}
		graph = append(graph, nextCol)
	}
	// recursively go through the entire solution, energize squares (or? create new symbols for energized walls?)
	height := len(graph)
	width := len(graph[0])
	var dfs func (x int, y int, dir string)
	fmt.Println(height)
	fmt.Println(width)
	dfs = func(x, y int, dir string) {
		// fmt.Printf("searching x %d y %d", x, y)
		// first, break out if we leave the area
		if x < 0 || x >= width || y < 0 || y >= height{
			return
		}
		// then, energize our next point
		point := graph[y][x]
		if strings.Contains(point.pastDir, dir){
			return
		}
		if point.pastDir == ""{
			ans++
		}
		point.pastDir += dir
		graph[y][x] = point
		// change direction based on rune
		switch point.char{
		case '|':
			if dir=="L" || dir=="R"{
				dfs(x, y+1, "D")
				dir = "U"
			}
		case '-':
			if dir=="U" || dir=="D"{
				dfs(x+1, y, "R")
				dir = "L"
			}
		case '\\':
			switch dir{
			case "U":
				dir = "L"
			case "D":
				dir = "R"
			case "L":
				dir = "U"
			case "R":
				dir = "D"
			}
		case '/':
			switch dir{
			case "U":
				dir = "R"
			case "D":
				dir = "L"
			case "L":
				dir = "D"
			case "R":
				dir = "U"
			}
		}		
		// recurse based on direction
		switch dir{
		case "U":
			y--
		case "D":
			y++
		case "L":
			x--
		case "R":
			x++
		default:
			fmt.Printf("invalid direction %s\n", dir)
		}
		dfs(x, y, dir)

	}
	
	// count up whenever you energize a square
	// after all directions have been exhausted finish up
	dfs (0, 0, "R")
	return ans
}