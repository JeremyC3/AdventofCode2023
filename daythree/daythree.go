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
	fmt.Println(Gear(str))
	// fmt.Println(BonusGear(str))
}

func Gear(f string) (ans int){
	// thoughts: go through the set twice, once to get the symbols twice to add the numbers
	// create a set of valids?
	adjParts :=  make(map[string]bool)
	for y, row:=range strings.Split(f, "\n"){
		runeRow := []rune(row)
		for x, char := range runeRow{
			// if it's not a number or period then it's a symbol woo
			if char != '.' && (char < '0' || char > '9'){
				adjParts[fmt.Sprintf("%d,%d", x+1, y)]=true
				adjParts[fmt.Sprintf("%d,%d", x+1, y-1)]=true
				adjParts[fmt.Sprintf("%d,%d", x+1, y+1)]=true
				adjParts[fmt.Sprintf("%d,%d", x-1, y+1)]=true
				adjParts[fmt.Sprintf("%d,%d", x-1, y)]=true
				adjParts[fmt.Sprintf("%d,%d", x-1, y-1)]=true
				adjParts[fmt.Sprintf("%d,%d", x, y-1)]=true
				adjParts[fmt.Sprintf("%d,%d", x, y+1)]=true
			}
		}
	}
	// after getting all the symbol valid locations, now we go through again to start making valid numbers
	for y, row:=range strings.Split(f, "\n"){
		runeRow := []rune(row)
		validPart := false
		tempSum := 0
		for x, char := range runeRow{
			// check if your character is a number	
			if (char >= '0' && char <= '9'){
				// if it's a number, add it to the tempSum and then check for validity
				tempSum = tempSum*10 + int(char-'0')
				fmt.Printf("adding %s from coordinates %d:%d with a sum of %d\n", string(char), x, y, tempSum)
				if i := adjParts[fmt.Sprintf("%d,%d", x, y)]; i == true && validPart == false{
					validPart = true
				}
			} else if tempSum >0{
				// if there's a tempSum and it's valid, add it to the answer. set tempSum to 0 regardless.
				if validPart == true{
					ans += tempSum
				}
				tempSum = 0
				validPart = false
			}
			// if the number is the last in the row you also need to check if it can be added
			if x == len(runeRow)-1 && validPart == true{
				ans += tempSum
			}
		}
	}
	return ans
}

func BonusGear(f string) (ans int){
	// thoughts: go through the set twice, once to get the symbols twice to add the numbers
	// create a set of valids?
	type partMap struct{
		Ratio int
		Adj int
	}
	adjParts :=  make(map[string]partMap)

	// make a struct for each valid part to map to
	// each part should have a running sum + adjacent count.
	
	
	for y, row:=range strings.Split(f, "\n"){
		for x, char := range row{
			// now we only need to add gears
			if char == '*'{
				adjParts[fmt.Sprintf("%d,%d", x, y)] = partMap{0, 0}
			}
		}
	}
	// after getting all the symbol valid locations, now we go through again to start making valid numbers
	for y, row:=range strings.Split(f, "\n"){
		runeRow := []rune(row)
		adjGears := make(map[string]bool)
		tempSum := 0
		for x, char := range runeRow{
			// check if your character is a number	
			if (char >= '0' && char <= '9'){	
				// if it's a number, add it to the tempSum and then check for validity
				tempSum = tempSum*10 + int(char-'0')
				fmt.Printf("adding %s from coordinates %d:%d with a sum of %d\n", string(char), x, y, tempSum)

				// check all adjacent spaces and see if we have any valid spaces.
				adjSpaces := [8]string{fmt.Sprintf("%d,%d", x+1, y+1),fmt.Sprintf("%d,%d", x+1, y) ,fmt.Sprintf("%d,%d", x+1, y-1), fmt.Sprintf("%d,%d", x, y+1),fmt.Sprintf("%d,%d", x, y-1), fmt.Sprintf("%d,%d", x-1, y+1), fmt.Sprintf("%d,%d", x-1, y), fmt.Sprintf("%d,%d", x-1, y-1)}
				for _, space := range adjSpaces{
					if _, ok:= adjParts[space]; ok{
						adjGears[space] = true
					}
				}
			} 
			// if we have values in our tempSum and either see an non-numeric character or reach the end of the row, process it
			if tempSum >0 && (char < '0' || char > '9' || x ==len(runeRow)-1){
				for gear := range adjGears {
					// first add 1 to the gear ratio
					tempAdj := adjParts[gear].Adj+1
					// if there's no adj gears, add the tempSum to the ratio.
					var newGear partMap
					switch tempAdj{
					case 1:
						newGear = partMap{tempSum, tempAdj}
					case 2:
						// if this Adj is now 2, get the product of tempSum and ratio and then update answer
						newRatio := adjParts[gear].Ratio * tempSum
						ans += newRatio
						newGear = partMap{newRatio, tempAdj}
					case 3:
						// now that there's 3 digits the previous ratio needs to be removed.
						ans -= adjParts[gear].Ratio
						newGear = partMap{0, tempAdj}
					default:
						// if the value's greater than e you don't need to do anything
						newGear = partMap{0, tempAdj}
					}
					adjParts[gear]=newGear
					
				}
				// clean up the tempSum as well as gear list
				tempSum = 0
				clear(adjGears)
			}
		}
	}
	return ans
}