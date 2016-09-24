package write

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/xDarkicex/writer/config"
)

//Note is the function that takes all notes
func Note() {
	// Open the file for read and write (O_RDRW), append to it if it has
	// content, create it if it does not exit, use 0666 for permissions
	// on creation.

	file, _ := os.OpenFile(config.File, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	// Close the file when the surrounding function exists
	defer file.Close()

	// Read old content
	current, _ := ioutil.ReadAll(file)
	if current != nil {
		fmt.Println("Currently Writing notes to " + config.File)
	}
	fmt.Print("Enter New Note: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		fmt.Print("Enter New Note: ")
		userInput := scanner.Text()
		if userInput == ":wq" {
			fmt.Println(" Saving...")
			fmt.Println("Quiting now")
			break
		}
		// Now write the input back to file text.txt
		file.WriteString(userInput + (time.Now().Format(time.Kitchen)) + "\n")
	}

}
