package main

import (
	"os"
	"testing"
)

func TestWasteland( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Wasteland(string(file)); y != 6440{
		t.Fatalf("expected value 6440, received %v", y)
	}
}

func TestBonusWasteland(t *testing.T){
	file, err := os.ReadFile("bonustestinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := BonusWasteland(string(file)); y != 6{
		t.Fatalf("expected value 6, received %v", y)
	}
}