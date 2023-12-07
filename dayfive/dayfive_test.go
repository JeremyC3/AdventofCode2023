package main

import (
	"os"
	"testing"
)

func TestSeed( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Seed(string(file)); y != 35{
		t.Fatalf("expected value 35, received %v", y)
	}
}

// func TestBonusSeed(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusSeed(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }