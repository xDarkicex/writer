package read

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xDarkicex/writer/config"
)

//Note This is the fucntion to read notes
func Note() {
	file, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Print all not4es to console
	notes, err := ioutil.ReadAll(file)
	fmt.Print(string(notes))
}
