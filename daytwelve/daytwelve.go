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
	fmt.Println(Spring(str))
	// fmt.Println(BonusSpring(str))
}

func Factorial(n int)(result int) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func GetWorking(s string, k int, n int )[]string{
	// break out if we put in too many #'s.
	if k<0 || n <0{
		return []string{}
	}
	if k==0{
		// if we've run out of #'s to place just go ahead and put in all .'s
		return []string{strings.ReplaceAll(s, "?", ".")}
	}
	// if we only have springs left just return them all
	if n==k{
		return []string{strings.ReplaceAll(s, "?", "#")}
	}
	// return string if:
	// we have no ?'s left to return (k=0)
	addSpring := strings.Replace(s, "?", "#", 1) 
	noSpring := strings.Replace(s, "?", ".", 1)
	return append(GetWorking(addSpring, k-1, n-1), GetWorking(noSpring, k, n-1)...)
}

func Spring(f string) (ans int){
	for _, row := range strings.Split(f, "\n"){
		fmt.Printf("previous row had %d correct \n", ans)
		records, notation, found  := strings.Cut(row, " ")
		if found == false{
			fmt.Println("string not found!")
			continue
		} 
		n := 0
		k := 0
		// push the numbers into an array
		// figure out how many question marks we have and how many 
		intNotation := []int{}
		for _, num := range strings.Split(notation, ","){
			if i, ok := strconv.Atoi(num); ok == nil {
				k += i
				intNotation = append(intNotation, i)
			}
		}
		for _, char := range records{
			if char == '#'{
				k--
			} else if char == '?'{
				n++
			}
		}
		// build out an arr of all possible strings
		potentialStr := GetWorking(records, k, n)
		ValidateStrings:
		for _, filledStr := range potentialStr{
			adjSprings := 0
			// validate all possible strings
			notationInd := 0
			// add all valid answers to the list
			valid := true
			for _, char := range filledStr{
				if char == '#'{
					adjSprings++
				} else if char == '.' {
					// if we just finished up a chunk of springs, check if it matches the next value
					if adjSprings>0{
						if notationInd < len(intNotation) && intNotation[notationInd]== adjSprings{
							notationInd++
						} else{
							valid = false
							continue ValidateStrings
						}
					} 
					adjSprings = 0
				} else {
					fmt.Println("invalid rune!!!")
					valid = false
					continue ValidateStrings
				}
			}
			if valid && notationInd >= len(intNotation)-1{
				ans++
			}
		}


		// bottomtext := Factorial(n)/(Factorial(k)*Factorial(n-k))
		// ans += bottomtext
		// fmt.Printf("this row has %d springs to fill into %d spots, and needs %d reps to get the answer.\n", k, n, bottomtext)
	}
	return ans
}