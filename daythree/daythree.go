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