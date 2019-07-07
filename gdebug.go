package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var recordList []interface{}

func check (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func htmlCtrl(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("debug.html");check(err)
	html := string(htmlByte)
	w.Header().Set("Content-Type", "text/html")
	_, err = fmt.Fprintf(w, html);check(err)
}
func debugCtrl(w http.ResponseWriter, r *http.Request) {
	var dataVO = map[string]interface{}{
		"type": "pass",
		"data": recordList,
	}
	//log.Print(recordList)
	dataJO, _ := json.Marshal(dataVO)
	_, _ = fmt.Fprintf(w, string(dataJO))
}
func clearDebugCtrl(w http.ResponseWriter, r *http.Request) {
	var dataVO = map[string]interface{}{
		"type": "pass",
		"data": struct{}{},
	}
	recordList = recordList[:len(recordList)-1]
	dataJO, _ := json.Marshal(dataVO)
	_, _ = fmt.Fprintf(w, string(dataJO))
}
func addDebugCtrl(w http.ResponseWriter, r *http.Request) {
	var dataVO = map[string]interface{}{
		"type": "pass",
		"data": "",
	}
	r.ParseForm() // 解析请求
	reqJSON := r.PostFormValue("json")
	var JSONVO  []interface{}
	err := json.Unmarshal([]byte(reqJSON), &JSONVO)
	check(err)
	recordList = append(recordList, JSONVO...)
	dataJO, _ := json.Marshal(dataVO)
	_, _ = fmt.Fprintf(w, string(dataJO))
}
func startServer () {
	http.HandleFunc("/", htmlCtrl)
	http.HandleFunc("/debug", debugCtrl)
	http.HandleFunc("/debug/add", addDebugCtrl)
	http.HandleFunc("/debug/clear", clearDebugCtrl)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Print("listen success")
	}
}

func main () {
	startServer()
}
