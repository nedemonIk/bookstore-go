package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter your grade")
	reader := bufio.NewReader(os.Stdin)   //for reading from the keabord
	input, err := reader.ReadString('\n') //reading from the user before the Enter
	if err != nil {                       // if error, alert a message and stop running
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)            //delete the symbol of the new row
	grade, err := strconv.ParseFloat(input, 64) // parsing the input to float64 and give
	// to grade
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failed"
	}
	fmt.Println("a grade of", grade, "is", status)
}
