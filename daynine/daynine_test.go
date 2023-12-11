package main

import (
	"os"
	"testing"
)

func TestOasis( t *testing.T){
	file, err := os.ReadFile("testinput")
	if err!=nil{
		t.Fatalf("error reading file")
	}

	if y := Oasis(string(file)); y != 114{
		t.Fatalf("expected value 114, received %v", y)
	}
}
func TestTriangle(t *testing.T){
	if y := Triangle([]int{1, 3, 6, 10, 15, 21}); y!= 28{
		t.Fatalf("expected value 28, received %v", y)
	}
}
// func TestBonusOasis(t *testing.T){
// 	file, err := os.ReadFile("bonustestinput")
// 	if err!=nil{
// 		t.Fatalf("error reading file")
// 	}

// 	if y := BonusOasis(string(file)); y != 30{
// 		t.Fatalf("expected value 30, received %v", y)
// 	}
// }