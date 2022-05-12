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

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":80", nil)

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
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
