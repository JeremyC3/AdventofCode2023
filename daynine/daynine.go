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
	fmt.Println(Oasis(str))
	// fmt.Println(BonusOasis(str))
}

func Triangle(hist []int) int{
	// break out when we have a 0?
	identical := true
	// make a new array and return the latest value minus
	newHist := []int{}
	for i:= 1; i<len(hist); i++{
		diff := hist[i]-hist[i-1]
		if diff != 0{
			identical = false
		}
		newHist = append(newHist, diff)
	}
	if identical == true{
		return hist[0]
	}

	return hist[len(hist)-1]+Triangle(newHist)
}

func Oasis(f string)(ans int){

	for _, row := range strings.Split(f, "\n"){
		// for each row, start only with the first 2 and keep increasing the triangle
		// until we have to leave?
		
		intRow := []int{}
		for key, num := range strings.Fields(row){
			if len(num) > 15{
				fmt.Printf("potential issue on line %d with string %s", key, num)
			}
			i, err := strconv.Atoi(num)
			if err != nil{
				fmt.Println("issue parsing num")
			} else{
				intRow = append(intRow, i)
			}
		}
		if len(intRow)>0{
			newAns := Triangle(intRow)
			fmt.Println(newAns)
			ans += newAns
		} 
	}
	return ans
}