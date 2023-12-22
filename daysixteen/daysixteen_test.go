package main

import (
	"os"
	"testing"
)

func TestLava( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Lava(string(file)); y != 46{
		t.Fatalf("expected value 46, received %v", y)
	}
}

// func TestBonusLava(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusLava(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }