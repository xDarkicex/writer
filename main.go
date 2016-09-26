package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/xDarkicex/writer/config"
	"github.com/xDarkicex/writer/read"
	"github.com/xDarkicex/writer/write"
)

func main() {
	fmt.Println("Go Pad Version: " + config.Version)
	config.Key()
	Password := config.Password
	config.GetAuthor()
	Author := config.Author
	if Password == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter New Password: ")
		for scanner.Scan() {
			Password = scanner.Text()
			fmt.Println(Password)
			config.PassAuth(Password)
		}
	} else if Author == "" {
		config.SetAuthor()
		main()
	} else {

		fmt.Println("Welcome " + Author)
		fmt.Println("Do you want to Write or Read a Note")
		fmt.Println("options(Read, Write, Exit)")
		fmt.Print("Enter Option: ")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			option := scanner.Text()
			if option == "write" || option == "Write" {
				write.Note()
				main()
			} else if option == "read" || option == "Read" {
				Auth()
			} else if option == "Exit" || option == "exit" {
				os.Exit(1)
				break
			}
		}
	}
}

//Auth Authenticates read access
func Auth() {

	if config.Session(config.IsLogged) == true {
		fmt.Println("Access Granted")
		read.Note()
		main()
	}
	fmt.Print("Enter Password: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Line := scanner.Text()

		fmt.Print("Enter Password: ")
		if err := bcrypt.CompareHashAndPassword([]byte(config.Password), []byte(Line)); err == nil {
			fmt.Println("Access Granted")
			Login := true
			config.Session(Login)
			read.Note()
			main()
		} else {
			fmt.Println("Password incorrect")
			fmt.Println("Unauthorized to read " + config.File)
			fmt.Printf("\n")
			fmt.Print("Enter Password: ")
			continue

		}
	}
}
