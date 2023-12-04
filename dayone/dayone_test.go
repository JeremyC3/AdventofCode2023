package main

import (
	"os"
	"testing"
)

func TestTrebuchet( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Trebuchet(string(file)); y != 142{
		t.Fatalf("expected value 142, received %v", y)
	}
}

func TestBonusTrebuchet(t *testing.T){
	file, err := os.ReadFile("bonustestinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := BonusTrebuchet(string(file)); y != 281{
		t.Fatalf("expected value 281, received %v", y)
	}
}