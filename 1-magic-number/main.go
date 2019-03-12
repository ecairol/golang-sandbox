package main

import (
	"fmt" 
	"bufio" 
	"os"
	"math/rand"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	
	// ReadString() reads until the first occurrence of \n
	var name, err = reader.ReadString('\n') 
	if err != nil {
		fmt.Println("Error." + err.Error())	
	}

	// Remove the last \n character from the string
	// Split the array from 0 to length-1
	if lastValidChar := len(name)-1; lastValidChar >= 0 {
		name = name[:lastValidChar]
	}

	// Initialize global Source 
	// Using Seed() function to initialize the default Source, since different behavior is required for each run.
	rand.Seed(time.Now().UnixNano())
	magicNumber := rand.Intn(99)
	fmt.Printf("%s, you should buy the lottery ticket %d \n", name, magicNumber)
}
