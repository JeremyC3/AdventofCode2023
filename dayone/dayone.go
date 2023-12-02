package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.ReadFile("input")
	if err != nil{
		fmt.Print(err)
	}
	str := string(file)
	fmt.Println(BonusTrebuchet(str))
}

func Trebuchet(f string) (ans int){
	// take the file and split it based on each new line
	// keep a running sum
	n := "1234567890"

	// iterate through each string in the array
	for _, v := range strings.Split(f, "\n"){
		// fmt.Printf(v)
		// record first and last number, turn it into a 2-digit # and then add to running sum
		first := strings.IndexAny(v, n)
		last := strings.LastIndexAny(v, n)
		if (first == -1 || last == -1){
			continue
		}
		sum := (string(v[first]) + string(v[last]))
		num, err := strconv.Atoi(sum)
		if (err != nil){
			fmt.Println(("oopsie"))
			continue
		}
		ans = num + ans
	}
	return ans
}

func BonusTrebuchet (f string) (ans int){
	// replace first, then take the string and finish everything.
	for k, v := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		f = strings.Replace(f, v, v+strconv.Itoa(k)+v , -1)
	}
	fmt.Println(f)
	return Trebuchet(f)
}