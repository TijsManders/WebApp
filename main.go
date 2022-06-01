package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
	SecureIT       []AlarmData
)

func main() {

	http.HandleFunc("/", RadioButtons)
	http.HandleFunc("/api", OntvangAPI)
	fmt.Println("Het dashboard is ")
	http.ListenAndServe(":5000", nil)
}

func OntvangAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API pagina")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var aData AlarmData
	json.Unmarshal(reqBody, &aData)
	SecureIT = append(SecureIT, aData)
	json.NewEncoder(w).Encode(aData)
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
