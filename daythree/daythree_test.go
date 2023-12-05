package main

import (
	"os"
	"testing"
)

func TestCube( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Gear(string(file)); y != 4361{
		t.Fatalf("expected value 4361, received %v", y)
	}
}

// func TestBonusGear(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusGear(string(file)); y != 2286{
// 		t.Fatalf("expected value 2286, received %v", y)
// 	}
// }