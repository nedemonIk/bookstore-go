package main

import "fmt"

func main() {
	var a, b int
	isfound := false
	fmt.Scan(&a)
	for fmt.Scan(&b); a <= b; b-- {
		if b%7 == 0 {
			isfound = true
			break
		} else {
			isfound = false
		}
	}
	if isfound {
		fmt.Print(b)
	} else {
		fmt.Print("NO")
	}
}
