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
	fmt.Println(Mirror(str))
	// fmt.Println(BonusMirror(str))
}

func Mirror(f string) (ans int){
	islands := strings.Split(f, "\n\n")
	IslandCheck:
	for key, island := range islands{
		fmt.Printf("Checking island #%d", key)
		islandMatrix := [][]rune{}
		prevRow := ""
		potentialCols := []int{}
		potentialRows := []int{}
		for key, row := range strings.Split(island, "\n"){
			if row == ""{
				continue
			}
			if prevRow == row{
				potentialRows = append(potentialRows, key-1)
				// start the comparisons now
			}
			prevRow = row
			islandMatrix = append(islandMatrix, []rune(row))
		}
		height := len(islandMatrix)
		width := len(islandMatrix[0])
		for i:=0; i<width-1; i++{
			if islandMatrix[0][i]==islandMatrix[0][i+1]{
				potentialCols = append(potentialCols, i)
			}
		}
		fmt.Println(potentialCols)
		fmt.Println(potentialRows)
		ValidateColumns:
		for _, colStart := range potentialCols{
			fmt.Printf("Checking col %d", colStart)
			colCheck := colStart
			mirrorCol := colCheck+1
			for(mirrorCol < width && colCheck >= 0){
				for i:=0; i<height; i++{
					if islandMatrix[i][colCheck]!=islandMatrix[i][mirrorCol]{
						continue ValidateColumns
					}
				}
				mirrorCol++
				colCheck--
			}
			ans += colStart+1
			continue IslandCheck
		}
		ValidateRows:
		for _, rowStart := range potentialRows{
			fmt.Printf("Checking row %d", rowStart)
			rowCheck:= rowStart
			mirrorRow := rowCheck+1
			// check until either rowCheck hits 0 or mirroRow hits the max
			for(mirrorRow < height && rowCheck>=0){
				if string(islandMatrix[rowCheck])!=string(islandMatrix[mirrorRow]){
					continue ValidateRows
				}
				mirrorRow++
				rowCheck--
			}
			// if you manage to break out of the for loop it's because everything is equal

			ans += (rowStart+1)*100
			continue IslandCheck
		}
	}
	fmt.Println("done")
	return ans
}