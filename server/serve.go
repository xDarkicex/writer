package server

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/GO-CLASS/lazy"
	"github.com/xDarkicex/writer/config"
)

var goHTML *template.Template
var goHTML2 *template.Template

func init() {

	goHTML = template.Must(template.ParseFiles("server/index.gohtml"))
	goHTML2 = template.Must(template.ParseFiles("server/error.gohtml"))
}

//LineItem in note
type LineItem struct {
	Author string
	Note   string
	Date   string
	Time   string
}

func handler(res http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.RequestURI)
	if len(req.RequestURI) >= 1 {
		if req.Method == "POST" {
			s.Say("New note: ", req.FormValue("AddNote"))

			file, _ := os.OpenFile(config.File, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
			// Close the file when the surrounding function exists
			defer file.Close()
			// Now write the input back to file
			file.WriteString(config.Author + ", " + req.FormValue("AddNote") + ", " + (time.Now().Format("Mon Jan 2 2006")) + ", " + (time.Now().Format(time.Kitchen)) + "\n")

		}
		if req.RequestURI == "/error" {
			s.Say("error URI in handler")
			errorPage(res, req)
		}
		if req.RequestURI == "/" {
			root(res, req)
		}
	} else {
		//
		s.Say(" --> Unhandled request at: ", req.RequestURI, req.Method, " <-- Request")
		// fmt.Println()
	}
}

//Fs file server exportablw
var Fs = http.FileServer(http.Dir("server/public"))

//HandleCSSRequest this is for css requests
func HandleCSSRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/css;")
}

//Serve Fucntion
func Serve() {
	s.Say("Listening on: 0.0.0.0:8080")
	s.Say("Status" + "Good")
	http.Handle("/server/public/", http.StripPrefix("/server/public/", Fs))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	// return router
}

func errorPage(res http.ResponseWriter, req *http.Request) {
	s.Say("At Error Page")
	reason := map[string]interface{}{
		"errorMessage": fmt.Sprint("Error in: " + config.File),
		"fixError":     ("To fix error in " + config.File + " please check /writer/" + config.File),
	}

	err := goHTML2.Execute(res, reason)
	if err != nil {
		log.Fatalln(err)
	}
}

func root(res http.ResponseWriter, req *http.Request) {

	s.Say(req.Method + " " + req.Host + " " + req.RequestURI)

	file, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	rows, err := rdr.ReadAll()
	if err != nil {
		s.Say("--> Error in: " + config.File + "<--")
		http.Redirect(res, req, "/error", 302)
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
