package main

import (
	"os"
	"testing"
)

func TestBeam( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Beam(string(file)); y != 136{
		t.Fatalf("expected value 136, received %v", y)
	}
}

// func TestBonusBeam(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusBeam(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }