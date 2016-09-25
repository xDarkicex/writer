package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xDarkicex/writer/config"
	"github.com/xDarkicex/writer/read"
	"github.com/xDarkicex/writer/write"
)

func main() {
	config.Key()
	Password := config.Password
	fmt.Println(Password)
	if Password == "" {
		fmt.Println("testing step 2")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter New Password: ")
		for scanner.Scan() {
			fmt.Println("Step 4")
			Password := scanner.Text()
			fmt.Println(Password)
			config.PassAuth(Password)
		}
	} else {
		fmt.Println("Testing step 3")
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
}

//Auth Authenticates read access
func Auth() {

	fmt.Print("Enter Password: ")
	// for scanner.Scan() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print("Enter Password: ")
		if line == config.Password {
			fmt.Println("Access Granted")
			read.Note()
			break
		} else {
			fmt.Println("Password incorrect")
			fmt.Println("Unauthorized to read " + config.File)
			fmt.Printf("\n")
			fmt.Print("Enter Password: ")
			continue
		}
	}
}
