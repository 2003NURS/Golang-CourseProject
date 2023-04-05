// "\app\space\home.go"
package space

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
	

	https.HandleFunc("/welcome/", Welcome)
	https.ListenAndServe("localhost:8000", nil)
}

// 	t, err :=template.ParseFiles("templates/welcome.html")
	
//  if err != nil {
//   fmt.Fprintf(w, err.Error())
//  }
//  t.Execute(w, "welcome")