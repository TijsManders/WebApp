package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
}

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.ListenAndServe(":80", nil)

}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request) {
	Title := "Alarm Activatie"
	MyRadioButtons := []RadioButton{
		RadioButton{}
	}


	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	value := r.Form.Get("Activatie")
	fmt.Println(value)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

// func main() {
// 	tmpl := template.Must(template.ParseFiles("index.html"))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		data := TodoPageData{
// 			PageTitle: "My TODO list",
// 			Todos: []Todo{
// 				{Title: "Task 1", Done: false},
// 				{Title: "Task 2", Done: true},
// 				{Title: "Task 3", Done: true},
// 			},
// 		}
// 		value := r.Form.Get("Activatie")
// 		fmt.Println(value)
// 		tmpl.Execute(w, data)
// 	})
// 	http.ListenAndServe(":80", nil)
// }
