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

	if y := Cube(string(file)); y != 8{
		t.Fatalf("expected value 8, received %v", y)
	}
}

func TestBonusCube(t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := BonusCube(string(file)); y != 2286{
		t.Fatalf("expected value 2286, received %v", y)
	}
}