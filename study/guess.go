package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	secretNum := rand.Intn(100) + 1
	fmt.Println("Try to guess the number.Enter yours!")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println("Now You have", 10-attempts, "attempts.")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < secretNum {
			fmt.Println("Your guess is too LOW.")
		} else if guess > secretNum {
			fmt.Println("Your guess is too High")
		} else {
			success = true
			fmt.Println("You won")
			break
		}
	}
	if !success {
		fmt.Println("Attempts ended.You lost.The right answer is", secretNum)
	}

}
