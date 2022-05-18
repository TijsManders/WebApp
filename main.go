package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type TodoPageData struct {
	AlarmActive bool
	AlarmAan    bool
}

var (
	ActivatieValue bool
	AlarmValue     bool
)

func main() {
	Handlerequests()
}

func Handlerequests() {
	http.HandleFunc("/", RadioButtons)
	http.HandleFunc("/", Get)
	http.ListenAndServe(":80", nil)
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("index.html"))
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		ActivatieV, err := strconv.ParseBool(r.Form.Get("Activatie"))
		if err != nil {
			log.Fatal(err)
		}
		ActivatieValue = ActivatieV
		fmt.Println(ActivatieValue, "Tijs")

		AlarmV, err := strconv.ParseBool(r.Form.Get("Status"))
		if err != nil {
			log.Fatal(err)
		}
		AlarmValue = AlarmV
		fmt.Println(AlarmValue, "Niet Tijs")

		data := TodoPageData{
			AlarmActive: ActivatieValue,
			AlarmAan:    AlarmValue,
		}
		tmpl.Execute(w, data)
	}
}

func RadioButtons(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		ActivatieV, err := strconv.ParseBool(r.Form.Get("Activatie"))
		if err != nil {
			log.Fatal(err)
		}
		ActivatieValue = ActivatieV
		fmt.Println(ActivatieValue, "Tijs")

		AlarmV, err := strconv.ParseBool(r.Form.Get("Status"))
		if err != nil {
			log.Fatal(err)
		}
		AlarmValue = AlarmV
		fmt.Println(AlarmValue, "Niet Tijs")
	}
}

// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"strconv"
// )

// type TodoPageData struct {
// 	AlarmActive bool
// 	AlarmAan    bool
// }

// func main() {
// 	tmpl := template.Must(template.ParseFiles("index.html"))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		var ActivatieValue bool
// 		var AlarmValue bool

// 		if r.Method == http.MethodPost {
// 			err := r.ParseForm()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			ActivatieV, err := strconv.ParseBool(r.Form.Get("Activatie"))
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			ActivatieValue = ActivatieV
// 			fmt.Println(ActivatieValue, "Tijs")

// 			AlarmV, err := strconv.ParseBool(r.Form.Get("Status"))
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			AlarmValue = AlarmV
// 			fmt.Println(AlarmValue, "Niet Tijs")

// 		}
// 		data := TodoPageData{
// 			AlarmActive: ActivatieValue,
// 			AlarmAan:    AlarmValue,
// 		}

// 		tmpl.Execute(w, data)
// 	})
// 	http.ListenAndServe(":80", nil)
// }
