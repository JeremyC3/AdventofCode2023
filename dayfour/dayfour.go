package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(Scratchcard(str))
	// fmt.Println(BonusScratchcard(str))
}

func Scratchcard(f string) (ans float64){
	// split each row into its own entry
	for _, entry := range strings.Split(f, "\n"){
		_, after, found := strings.Cut(entry, ":")
		if(found == false){
			continue
		}
		winningNums := make(map[string]bool)
		yourNums := false
		hits := -1.0
		for _, number := range strings.Fields(after){
			if number == "|"{
					yourNums = true
			} else if yourNums == false {
				winningNums[number]=true
			} else if win := winningNums[number]; win==true {
				hits +=1
			}
		}
		if hits >= 0 {
			ans += math.Pow(2 , hits)
		}
	}
	// split each entry into winning nums + numbers
	// count the number of entries you get right, give 2^n-1 pts
	// add all of them in at the end 
	return ans
}