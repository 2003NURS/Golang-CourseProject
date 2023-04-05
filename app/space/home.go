// "\app\space\home.go"
package space

import (
	"fmt"
	"net/http"
	"errors"
	"html/template"
	"runtime/internal/sys"
)

var data = ""

func Welcome(w http.ResponseWriter, r *http.Request){
	fmt.Println("Login running")
	t, err := template.ParseFiles("templates/welcome.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerRequest(){
	http.HandleFunc("/welcome/", Welcome)
	http.ListenAndServe("localhost:8000", nil)
}
