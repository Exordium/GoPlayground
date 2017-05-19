package main

import (
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
	correctAnswer := rand.Intn(10)
	done := false
	for !done {
		fmt.Print("Guess a number between 1 and 10: ")
		reader := bufio.NewReader(os.Stdin)
        var e error
		input, e := reader.ReadString('\n')
		if e != nil{
			fmt.Println("Definitely something stupid happened.")
		}
		num, e := strconv.Atoi(strings.TrimSpace(input))
    	if e != nil {
			fmt.Println(e)
		} else if num == correctAnswer {
			fmt.Println("YOU DID IT! :D  The number was", correctAnswer)
			done = !done
		} else {
			fmt.Println("Nope, you suck.")
		}
	}
}