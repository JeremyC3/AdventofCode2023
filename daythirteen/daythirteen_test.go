package main

import (
	"os"
	"testing"
)

func TestMirror( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Mirror(string(file)); y != 405{
		t.Fatalf("expected value 405, received %v", y)
	}
}

// func TestBonusMirror(t *testing.T){
// 	file, err := os.ReadFile("testinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusMirror(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }