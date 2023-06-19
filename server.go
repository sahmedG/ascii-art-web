package webart

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/index.html"))
}

type RESULT_ASCII_ART struct {
	Result   bool
	AsciiArt string
	Color    string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "../templates/pagenotfound.html")
			return
		} else {
			tpl, err := template.ParseFiles("../templates/index.html")
			if err != nil {
				fmt.Fprint(w, err.Error())
			}
			tpl.ExecuteTemplate(w, "homepage", "")
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "../templates/badrequest.html")
	}
}

func Asciiart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err:%v", err)
			return
		} else {

			value := r.FormValue("Text")
			Banner := r.FormValue("Banner")
			color := r.FormValue("color")
			if MapFont(Banner) == "" || value == "" {
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "../templates/iternalerror.html")
				return
			} else {
				text := PrintART(value, Banner)

				data := RESULT_ASCII_ART{true, text, color}
				//err := tpl.ExecuteTemplate(w, "homepage", data)
				err := tpl.ExecuteTemplate(w, "homepage", data)
				if err != nil {
					http.Error(w, err.Error(), 500)
					log.Fatalln(err)
				}
			}
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "../templates/badrequest.html")
	}
}