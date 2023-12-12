package main

import (
	"os"
	"testing"
)

func TestPipe( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Pipe(string(file)); y != 8{
		t.Fatalf("expected value 8, received %v", y)
	}
}

// func TestBonusPipe(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusPipe(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }