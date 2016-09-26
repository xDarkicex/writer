package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

//Version of note taker
var Version = "1.0.0"

//File to be converted to notes
var File = "note.gotxt"

//Author of note
var Author string

//SessionKey is for stored memory login
var SessionKey string

//IsLogged declartion
var IsLogged = false

//KeyLocal will store Private vars
var KeyLocal = "key.gotxt"

//AuthorLocal will store Private author
var AuthorLocal = "author.gotxt"

//Password Global declaration
var Password string

//Password is for reading file

//PassAuth location to password
func PassAuth(Password string) {

	file, err := os.OpenFile(KeyLocal, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("failed to openFile")
		log.Fatal(err)
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Somethign has gone terribly wrong..")
	}

	file.WriteString(string(hashedPass))
	defer file.Close()
}

//SetAuthor Function to set author
func SetAuthor() {
	file, err := os.OpenFile(AuthorLocal, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("failed to openFile")
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Set New Author: ")
	for scanner.Scan() {
		fmt.Print("Set New Author: ")
		Author := scanner.Text()
		file.WriteString(Author)
		defer file.Close()
		break
	}
}

//Key gets key value out of file
func Key() string {

	file, err := os.Open(KeyLocal)
	if err != nil {
		log.Fatal(err)
	}

	PassKey, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error getting Key")
	}
	Password = string(PassKey)
	defer file.Close()
	return Password
}

//GetAuthor gets author name from file
func GetAuthor() string {
	file, err := os.Open(AuthorLocal)
	if err != nil {
		//Custom Error handling attempting to repair problem
		fmt.Println(err)
		fmt.Println("Can not open author.txt")
	}
	Name, err := ioutil.ReadAll(file)
	if err != nil {
		//Attempt Recovery
		fmt.Println(err)
		fmt.Println("Auther name Not Found")
	}
	Author = string(Name)
	defer file.Close()

	return Author
}

//Session Data
func Session(Login bool) bool {
	IsLogged = Login
	return IsLogged
}
