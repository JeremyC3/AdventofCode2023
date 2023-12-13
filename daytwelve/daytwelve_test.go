package main

import (
	"fmt"
	"os"
	"testing"
)

func TestRecursion( t *testing.T){
	if i := GetWorking("????.######..#####.", 1, 4); i!= nil{
		fmt.Println(i)
		t.Fatalf("idk lol")
	}
}

func TestSpring( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Spring(string(file)); y != 21{
		t.Fatalf("expected value 21, received %v", y)
	}
}

// func TestBonusSpring(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusSpring(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }