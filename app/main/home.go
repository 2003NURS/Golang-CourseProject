// "\app\main\home.go"
package home
func Welcome(w http.ResponseWriter, r *http.Request){
	fmt.Println("Login running")
	t, err :=template.ParseFiles("templates/welcome.html")
	
 if err != nil {
  fmt.Fprintf(w, err.Error())
 }
 t.Execute(w, "welcome")
}

function HandlerRequest(){
	

	https.HandleFunc("/welcome/", Welcome)
	https.ListenAndServe("localhost:8000", nil)
}