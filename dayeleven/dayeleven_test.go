package main

import (
	"os"
	"testing"
)

func TestGalaxy( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Galaxy(string(file)); y != 374{
		t.Fatalf("expected value 374, received %v", y)
	}
}

func TestBonusGalaxy(t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := BonusGalaxy(string(file)); y != 8410{
		t.Fatalf("expected value 8410, received %v", y)
	}
}