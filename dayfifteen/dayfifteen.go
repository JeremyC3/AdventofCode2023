package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(Hash(str))
	// fmt.Println(BonusHash(str))
}

func Hash(f string)(ans int){
	for _, input := range strings.Split(f, ","){
		currVal := 0
		for ind := range input{
			currVal += int(input[ind])
			currVal *= 17
		}
		currVal = currVal%256
		ans += currVal
	}
	return ans
}
func BonusHash(f string)(ans int){
	// type LL struct{
	// 	name string
	// 	lens int
	// 	next *LL
	// }
	boxes := map[int][]string{}
	for _, input := range strings.Split(f, ","){
		currVal := 0
		for ind, char := range input{
			if char == '-'{
				// if we're a minus, go find the box and kill the element
				hash := currVal%256
				if i:= slices.Index(boxes[hash], input[:ind]); i>=0{
					
				} else{
					continue
				}
			} else if char == '='{
				hash := currVal%256
				if _, ok := boxes[hash]; ok{
					boxes[hash] = append(boxes[hash], input[:ind])
				} else{
					boxes[hash] = []string{input[:ind]}
				}

				// take the next element, convert it into a lens size and inspect it
			} else {
				currVal += int(input[ind])
				currVal *= 17
			}
		}
		currVal = currVal%256
		ans += currVal
	}
	return ans
}