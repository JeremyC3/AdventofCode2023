package main

import (
	"os"
	"testing"
)

func TestCamel( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Camel(string(file)); y != 6440{
		t.Fatalf("expected value 6440, received %v", y)
	}
}

// func TestBonusCamel(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusCamel(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }