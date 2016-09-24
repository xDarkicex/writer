package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// Open the file for read and write (O_RDRW), append to it if it has
	// content, create it if it does not exit, use 0666 for permissions
	// on creation.

	file, _ := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	// Close the file when the surrounding function exists
	defer file.Close()

	// Read old content
	current, _ := ioutil.ReadAll(file)

	// Do something with that old content, for example, print it
	fmt.Println(string(current))

	// Standard input from keyboard
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	// Now write the input back to file text.txt
	file.WriteString(userInput + "  " + (time.Now().Format(time.Kitchen)) + "\n")
}
