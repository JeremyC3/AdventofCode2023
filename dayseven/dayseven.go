package main

import (
	"cmp"
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
	fmt.Println(Camel(str))
	// fmt.Println(BonusCamel(str))
}

func Camel(f string)(ans int){
	// make a slice of structs to re-parse the info
	type hand struct{
		Cards []rune
		Bet int
		HandPow int
	}
	var handSlice []hand
	// split the entire thing based on newlines
	for rowNum, row :=range strings.Split(f, "\n"){
		cards, bet, didCut := strings.Cut(row, " ")
		betInt, err := strconv.Atoi(bet)
		if didCut && err == nil{
			handMap := make(map[rune]int)
			handPts := 0
			for _, rune := range cards{
				//count the number of hands
				handMap[rune]=handMap[rune]+1
				switch handMap[rune] {
				case 2:
					handPts += 1
				case 3:
					handPts += 2
				case 4:
					handPts += 4
				case 5:
					handPts += 6
				} 
			}
			// rate the hand based on the number of similars.
			// high = 0, pair = 1, twopair = 2, 3of = 3, fullhouse = 4, 4of = 5, 5 of = 6 
			newHand := hand{[]rune(cards), betInt, handPts}
			handSlice = append(handSlice, newHand)
		} else {
			fmt.Printf("unable to parse row %d", rowNum)
		}
	}
	cardRank := map[rune]int{'A':14, 'K':13, 'Q':12, 'J':11, 'T':10, '9':9,'8': 8, '7':7, '6':6, '5':5, '4':4,'3': 3, '2': 2}
	slices.SortFunc(handSlice,
	 func(a, b hand) int {
		if i:= cmp.Compare(a.HandPow, b.HandPow); i==0{
			for key := range a.Cards{
				aRune := cardRank[a.Cards[key]]
				bRune := cardRank[b.Cards[key]]
				if j:= cmp.Compare(aRune, bRune);j!=0{
					return j
				}
			}
			fmt.Println("same hand somehow weird")
			return 0
		} else {
			return i
		}
	 })
	// create helper function to compare two hands
	// sort based on higher hand value
	// reduce entire slice back down
	fmt.Println(handSlice)
	for mult := range handSlice{
		ans += handSlice[mult].Bet * (mult+1)
	} 
	return ans
}