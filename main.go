package main

import "fmt"

func main() {
	names := []string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	rangeOne := names[1:3]
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(names[len(names)-1])
}
