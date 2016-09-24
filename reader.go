package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xDarkicex/writer/config"
)

func main() {
	file, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Print all not4es to console
	notes, err := ioutil.ReadAll(file)
	fmt.Print(string(notes))
}
