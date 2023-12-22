package main

import (
	"fmt"
	"math"
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
	fmt.Println(Crucible(str))
	// fmt.Println(BonusCrucible(str))
}
// take in a pointer to the arr and also new coord + directions?

func Crucible(f string)(ans int){
	// first, make an array of all of the values
	// fmt.Println("you're watching disney channel")
	graph := [][]int{}
	for _, row := range strings.Split(f, "\n"){
		if row == ""{
			continue
		}
		newRow := []int{}
		for _, char := range row{
			newRow = append(newRow, int(char-'0'))
		}
		graph = append(graph, newRow)
	}
	height := len(graph)
	width := len(graph[0])
	complete := map[string]bool{}
	// make a set of visited locations
	// make a set of non-infinite locations?
	type block struct{
		minHeat int
		dI int
		dJ int
	}
	incomplete := map[string]block{"0,0":block{0, 0, 0}}
	// starting at 0,0 : explore all paths next to you (UDLR)
	i := 0
	j := 0
	dI := 0
	dJ := 0
	currStr := fmt.Sprintf("%d,%d", i, j)
	// break out if both i == height and j == width, so don't break as long as either one is false
	for(len(incomplete)>0 &&  (i!= height-1 || j != width-1) ){
		
		// current nodes need to be thrown into the complete section
		currStr = fmt.Sprintf("%d,%d", i, j)
		currHeat := incomplete[currStr].minHeat
		delete(incomplete, currStr)
		complete[currStr] = true
		// update nodes cardinally surrounding current node
		for _, vals := range [][]int{{i+1, j}, {i-1, j}, {i, j-1},{i, j+1} } {
			tempi := vals[0]
			tempDi := i-tempi
			tempj := vals[1]
			tempDj := j-tempj
			tempStr := fmt.Sprintf("%d,%d", tempi, tempj)
			_, isComplete := complete[tempStr]
			// get rid of invalid values, or completed values
			if ( isComplete|| tempi <0 || tempi >= height || tempj<0 || tempj >= width || tempDi+dI > 3 || tempDi+dI < -3 || tempDj+dJ > 3 || tempDj+dJ < -3 ){
				continue
			}
			tempHeat := graph[tempi][tempj]
			// get the minimum if the value exists, or just assign it to the current val
			if old, ok := incomplete[tempStr]; ok{
				if old.minHeat > currHeat + tempHeat{
					incomplete[tempStr] = block{currHeat+ tempHeat, dI + tempDi, dJ+ tempDj}
				}
			} else {
				incomplete[tempStr] = block{currHeat+ tempHeat, dI + tempDi, dJ+ tempDj}

			}
		}
		// after getting all surrounding nodes assessed, find the lowest value and assign it to our new value
		
		minVal := int(math.Inf(1))
		// fmt.Println("incomplete values")
		for key, newBlock := range incomplete {
			// fmt.Printf("%s: %d\n",key, val)
			val := newBlock.minHeat
			if val < minVal{
				minVal = val
				newCoords := strings.Split(key, ",")
				newi, erri := strconv.Atoi(newCoords[0])
				newj, errj := strconv.Atoi(newCoords[1])
				if errj!= nil || erri!=nil{
					panic(erri)
				}
				i = newi
				j = newj
				dI = newBlock.dI
				dJ = newBlock.dJ
			}
		}
		fmt.Printf("next spot will be i%d, j%d\n", i, j)
		
	}
	fmt.Println(height)
	fmt.Println(width)
	// fmt.Printf("%d,%d", i, j)
	// update each of the new locations with its net value to get there
	// update each direction if the new path is cheaper
	// go to the next lowest value
	return incomplete[fmt.Sprintf("%d,%d", i, j)].minHeat
	// return -1
}

