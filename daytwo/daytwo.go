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
	fmt.Println(Cube(str))
	fmt.Println(BonusCube(str))
}

func Cube(f string) (ans int){
	//split up the string based on newlines
	for k, v:= range strings.Split(f, "\n"){
		// for each game: split up each of the draws based on semicolons
		if len(v) ==0 {
			continue
		}
		if k<0{
			continue
		}
		tempGame := strings.Split(v, ":")
		games := tempGame[1]
		legal := true
		for _, draw := range strings.Split(games, ";"){
			// iterate through the string, buffering the numbers
			ballNum := 0
			for _, field := range strings.Fields(draw){
				// once you get a letter figure out if it's red blue or green
				if i, err := strconv.Atoi(field); err == nil {
					ballNum = i
				} else {
					switch string(field[0]) {
					case "r":
						if ballNum > 12{
							legal = false
						}
					case "g":
						if ballNum > 13{
							legal = false
						}
					case "b":
						if ballNum > 14{
							legal = false
						}
					}	
				}
				if legal == false{
					continue
				}
			}
			if legal == false{
				continue
			}
		}  
		if legal == true {
			ans += k+1
		}
	}
	return ans
}

func BonusCube(f string) (ans int){
	//split up the string based on newlines
	for k, v:= range strings.Split(f, "\n"){
		// for each game: split up each of the draws based on semicolons
		if len(v) ==0 {
			continue
		}
		if k<0{
			continue
		}
		tempGame := strings.Split(v, ":")
		games := tempGame[1]
		red := 0
		blue := 0
		green := 0
		for _, draw := range strings.Split(games, ";"){
			// iterate through the string, buffering the numbers
			ballNum := 0
			
			for _, field := range strings.Fields(draw){
				// once you get a letter figure out if it's red blue or green
				
				if i, err := strconv.Atoi(field); err == nil {
					ballNum = i
				} else {
					switch string(field[0]) {
					case "r":
						red = max(red, ballNum)
					case "g":
						green = max(green, ballNum)
					case "b":
						blue = max(blue, ballNum)
					}	
				}
			}
			}  
		ans += red*blue*green
	}
	return ans
}