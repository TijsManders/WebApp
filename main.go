package main

import (
	"bytes"
	"encoding/json"
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

type AlarmData struct {
	Activatie bool `json:"Activatie"`
	Alarm     bool `json:"Alarm"`
}

var (
	ActivatieValue bool
	AlarmValue     bool
)

func main() {
	http.HandleFunc("/", RadioButtons)
	http.ListenAndServe(":80", nil)
}

func StuurNaarAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Dit werkt ook")
		AlarmDataAPI := []AlarmData{
			{Activatie: ActivatieValue},
			{Alarm: AlarmValue},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(AlarmDataAPI)
		req, _ := http.NewRequest("POST", "localhost", payloadBuf)
		if req == nil {
			fmt.Println("hallo reg is nil")
		}
	}
}

func RadioButtons(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		ActivatieV, err := strconv.ParseBool(r.Form.Get("ActivatieStatus"))
		if err != nil {
			log.Fatal(err)
		}
		ActivatieValue = ActivatieV
		fmt.Println(ActivatieValue, "Tijs")

		AlarmV, err := strconv.ParseBool(r.Form.Get("AlarmStatus"))
		if err != nil {
			log.Fatal(err)
		}
		AlarmValue = AlarmV
		fmt.Println(AlarmValue, "Niet Tijs")
	}
	data := TodoPageData{
		AlarmActive: ActivatieValue,
		AlarmAan:    AlarmValue,
	}
	tmpl.Execute(w, data)
	StuurNaarAPI(w, r)
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
