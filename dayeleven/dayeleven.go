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
	// fmt.Println(Galaxy(str))
	fmt.Println(BonusGalaxy(str))
}

func Galaxy(f string)(ans int){
	starmap := strings.Split(f, "\n")
	starmap = starmap[:len(starmap)-1]
	width := len(starmap[0])
	height := len(starmap)
	fullCols := make([]bool, width)
	fullRows := make([]bool, height)
	type coord struct{
		X int
		Y int
	}
	galaxies := []coord{}
	// make a map of the galaxies
	// save all of the empty spots
	// push every map?
	for y, row := range starmap{
		for x, col := range row{
			if col == '#'{
				fmt.Printf("new Galaxy at %d, %d\n", x, y)
				galaxies = append(galaxies, coord{x, y})
				fullCols[x] = true
				fullRows[y] = true
			}
			
		}
	}
	emptyCols := []int{}
	emptyRows := []int{}
	for x, col := range fullCols{
		if col == false{
			emptyCols = append(emptyCols, x)
		}
	}
	for y, row := range fullRows{
		if row == false{
			emptyRows = append(emptyRows, y)
		}
	}
	for key, gal :=range  galaxies{
		yChange := 0
		xChange := 0
		for _, row :=range  emptyRows{
			if gal.Y > row{
				yChange++
			}
		}
		for _, col :=range  emptyCols{
			if gal.X > col{
				xChange++
			}
		}
		galaxies[key] = coord{gal.X+xChange, gal.Y+yChange}
	}
	fmt.Println(emptyCols)
	fmt.Println(emptyRows)
	fmt.Println(galaxies)
	for i, homeGal := range galaxies{
		for j:= i+1; j<len(galaxies); j++{
			destGal := galaxies[j]
			if diffX := homeGal.X-destGal.X; diffX>0{
				ans += diffX
			} else {
				ans -= diffX
			}
			if diffY := homeGal.Y-destGal.Y; diffY>0{
				ans += diffY
			} else {
				ans -= diffY
			}
		}
	}
	return ans
}

func BonusGalaxy(f string)(ans int){
	starmap := strings.Split(f, "\n")
	starmap = starmap[:len(starmap)-1]
	width := len(starmap[0])
	height := len(starmap)
	fullCols := make([]bool, width)
	fullRows := make([]bool, height)
	type coord struct{
		X int
		Y int
	}
	galaxies := []coord{}
	// make a map of the galaxies
	// save all of the empty spots
	// push every map?
	for y, row := range starmap{
		for x, col := range row{
			if col == '#'{
				galaxies = append(galaxies, coord{x, y})
				fullCols[x] = true
				fullRows[y] = true
			}
			
		}
	}
	emptyCols := []int{}
	emptyRows := []int{}
	for x, col := range fullCols{
		if col == false{
			emptyCols = append(emptyCols, x)
		}
	}
	for y, row := range fullRows{
		if row == false{
			emptyRows = append(emptyRows, y)
		}
	}
	for key, gal :=range  galaxies{
		yChange := 0
		xChange := 0
		for _, row :=range  emptyRows{
			if gal.Y > row{
				yChange += 999999
			}
		}
		for _, col :=range  emptyCols{
			if gal.X > col{
				xChange += 999999
			}
		}
		galaxies[key] = coord{gal.X+xChange, gal.Y+yChange}
	}
	// fmt.Println(galaxies)
	for i, homeGal := range galaxies{
		for j:= i+1; j<len(galaxies); j++{
			destGal := galaxies[j]
			if diffX := homeGal.X-destGal.X; diffX>0{
				ans += diffX
			} else {
				ans -= diffX
			}
			if diffY := homeGal.Y-destGal.Y; diffY>0{
				ans += diffY
			} else {
				ans -= diffY
			}
		}
	}
	return ans
}
