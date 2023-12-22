package main

import (
	"os"
	"testing"
)

func TestCrucible( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Crucible(string(file)); y != 102{
		t.Fatalf("expected value 102, received %v", y)
	}
}

// func TestBonusCrucible(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusCrucible(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }