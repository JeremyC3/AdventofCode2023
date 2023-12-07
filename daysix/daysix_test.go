package main

import (
	"os"
	"testing"
)

func TestBoat( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Boat(string(file)); y != 288{
		t.Fatalf("expected value 288, received %v", y)
	}
}

func TestBonusBoat(t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := BonusBoat(string(file)); y != 71503{
		t.Fatalf("expected value 71503, received %v", y)
	}
}