package main

import (
	"os"
	"testing"
)

func TestHash( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Hash(string(file)); y != 136{
		t.Fatalf("expected value 136, received %v", y)
	}
}

// func TestBonusHash(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusHash(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }