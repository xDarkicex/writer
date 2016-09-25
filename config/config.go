package config

import (
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

//Store will store Private vars
var Store = "Store.gotxt"

//Password is for reading file

//PassAuth location to password
func PassAuth(Password string) {

	file, err := os.OpenFile(Store, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("failed to openFile")
		log.Fatal(err)
	}
	fmt.Println("Step 5")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Somethign has gone terribly wrong..")
	}

	file.WriteString(string(hashedPass))
	defer file.Close()

}

//Key gets key value out of file
func Key() string {
	fmt.Println("Fired")
	file, err := os.Open(Store)
	if err != nil {
		log.Fatal(err)
	}

	PassKey, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error getting Key")
	}
	Password := string(PassKey)
	defer file.Close()
	fmt.Println(Password)
	return Password
	//Password is pass

}
