package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(Seed(str))
	// fmt.Println(BonusSeed(str))
}

func Seed(f string)(ans int){
	// slowest way first: check each seed for each new map line
	var seeds []int
	// guard to prevent double-transformation of a seed
	usedSeeds := make(map[int]bool)
	for key, line := range strings.Split(f, "\n"){
		if key==0{
			for _, seed := range strings.Fields(line){
				if i, err:= strconv.Atoi(seed); err==nil{
					seeds = append(seeds, i)
				}
			}
		} else if len(line)==0{
			continue
		} else if strings.HasSuffix(line, "map:"){
			//at each new map line, reset the seed list of used seeds
			clear(usedSeeds)
			continue
		} else {
			// first take all 3 nums and split them up
			mapArr := strings.Fields(line)
			mapS, errSource:= strconv.Atoi(mapArr[1])
			mapR, errRange := strconv.Atoi(mapArr[2])
			mapD, errDest := strconv.Atoi(mapArr[0])
			if errSource != nil || errRange != nil || errDest != nil{
				fmt.Println("oopsies")
				break
			}
			// fmt.Printf("source %d, range %d, dest %d\n", mapS, mapR, mapD)
			for seedKey, seedVal := range seeds{
				if _, i := usedSeeds[seedKey]; i!=true && seedVal >= mapS && seedVal < (mapS+mapR){
					seeds[seedKey] += mapD-mapS
					usedSeeds[seedKey]=true
					// fmt.Printf("converting seed %d at key %d to %d with line %d\n", seedVal,seedKey, seeds[seedKey], key)
				}
			}
		}
	}

	return slices.Min(seeds)
}