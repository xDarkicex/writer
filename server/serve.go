package server

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/writer/config"
)

var goHTML *template.Template

func init() {

	goHTML = template.Must(template.ParseFiles("server/index.gohtml"))
}

//LineItem in note
type LineItem struct {
	Author string
	Note   string
	Date   string
	Time   string
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI)
	if len(req.RequestURI) == 1 {
		root(res)
	} else {
		//
	}
}

//Fs file server exportablw
var Fs = http.FileServer(http.Dir("server/public"))

//HandleCSSRequest this is for css requests
func HandleCSSRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/css;")
}

//Web Add Notes Via Web interface
func Web(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("Did I fire?")
	UserInput := req.FormValue("AddNote")
	file, _ := os.OpenFile(config.File, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	// Close the file when the surrounding function exists
	defer file.Close()

	// Read old content
	current, _ := ioutil.ReadAll(file)
	if current != nil {
		fmt.Println("Currently Writing notes to " + config.File)
	}

	// Now write the input back to file
	file.WriteString(config.Author + ", " + UserInput + ", " + (time.Now().Format("Mon Jan 2 2006")) + ", " + (time.Now().Format(time.Kitchen)))

}

//Serve Fucntion
func Serve() *httprouter.Router {
	router := httprouter.New()
	http.Handle("/server/public/", http.StripPrefix("/server/public/", Fs))
	http.HandleFunc("/", handler)
	router.GET("/server/public/", HandleCSSRequest)
	router.POST("/note", Web)
	http.ListenAndServe(":8080", nil)
	return router
}
func root(res http.ResponseWriter) {

	file, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	LineItems := make([]LineItem, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}

		Author := row[0]
		Note := row[1]
		Date := row[2]
		Time := row[3]

		LineItems = append(LineItems, LineItem{
			Author: Author,
			Note:   Note,
			Date:   Date,
			Time:   Time,
		})
	}
	err2 := goHTML.Execute(res, LineItems)
	if err2 != nil {
		log.Fatalln(err)
	}
}
