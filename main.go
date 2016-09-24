package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xDarkicex/writer/config"
	"github.com/xDarkicex/writer/read"
	"github.com/xDarkicex/writer/write"
)

//Auth Authenticates read access
func Auth() {
	var Count = 0
	if Count < 3 {
		fmt.Print("Enter Password: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			line := scanner.Text()
			fmt.Print("Enter Password: ")
			if line == config.Password {
				fmt.Println("Access Granted")
				read.Note()
			} else {

				fmt.Println("Password incorrect")
				fmt.Println("Unauthorized to read " + config.File + "Remaining Atttempts " + string(Count))
			}
		}
		Count = Count + 1
	}
}

func main() {
	fmt.Println("Loading Gopad")
	fmt.Println("Do you want to Write or Read a Note")
	fmt.Println("options(Read, Write)")
	fmt.Print("Enter Option: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		option := scanner.Text()
		if option == "write" || option == "Write" {
			write.Note()
		} else if option == "read" || option == "Read" {
			Auth()
		} else {
			break
		}
	}
}
