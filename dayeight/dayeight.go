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
	fmt.Println(Wasteland(str))
	fmt.Println(BonusWasteland(str))
}

func Wasteland(f string)(ans int){
	type node struct{
		Left string
		Right string
	}
	nodeMap := make(map[string]node)
	//first make a map of all the new locations, then
	var instructions []rune
	steps := 0
	for key, row := range strings.Split(f, "\n"){
		runeRow := []rune(row)
		if key==0{
			instructions = runeRow
		} else if len(runeRow)>1{
			newNode := node{string(runeRow[7:10]), string(runeRow[12:15])}
			nodeMap[string(runeRow[0:3])] = newNode
		}
	}
	currNode := "AAA"
	for (currNode != "ZZZ"){
		direction := instructions[steps % len(instructions)]
		steps++
		if direction == 'L'{
			currNode = nodeMap[currNode].Left
		} else if direction == 'R'{
			currNode = nodeMap[currNode].Right
		} else{
			fmt.Printf("error during loop at step %d", steps)
			break;
		}
	}
	return steps
}

func EndswithZ(arr []string, steps int)(bool){
	for _, nodeStr := range arr{
		if strings.HasSuffix(nodeStr, "Z")==false {
			// fmt.Printf("%s doesn't end with Z", nodeStr)
			return false
		} else{
			
		}
	}
	fmt.Println("returning true!!")
	return true
}

func BonusWasteland(f string)(ans int){
	type node struct{
		Left string
		Right string
	}
	nodeMap := make(map[string]node)
	//first make a map of all the new locations, then
	var instructions []rune
	steps := 0
	currNodes := []string{}
	stepSlice := []int{}
	for key, row := range strings.Split(f, "\n"){
		runeRow := []rune(row)
		if key==0{
			instructions = runeRow
		} else if len(runeRow)>1{
			newNode := node{string(runeRow[7:10]), string(runeRow[12:15])}
			nodeMap[string(runeRow[0:3])] = newNode
			if runeRow[2] == 'A'{
				currNodes = append(currNodes, string(runeRow[0:3]))
			}
		}
	}
	// currNode := "AAA"
	// make a checker function to see if all the strings end with Z
	for (EndswithZ(currNodes, steps)==false){
		direction := instructions[steps % len(instructions)]
		steps++
		for ind, nextNode := range currNodes{
			if direction == 'L'{
			currNodes[ind] = nodeMap[nextNode].Left
		} else if direction == 'R'{
			currNodes[ind] = nodeMap[nextNode].Right
		} else{
			fmt.Printf("error during loop at step %d", steps)
			break;
		}
		// if we've just finished then we remove the elements that are finished
		}
	}
	return steps
}