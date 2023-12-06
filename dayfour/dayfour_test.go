package main

import (
	"os"
	"testing"
)

func TestScratchcard( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Scratchcard(string(file)); y != 13{
		t.Fatalf("expected value 13, received %v", y)
	}
}

// func TestBonusScratchcard(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusScratchcard(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }