
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", manageIndex)
	http.HandleFunc("/login", manageLogin)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func manageIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters and body
	///formPrm := r.Form
	fmt.Fprintln(w, "Hello HTTP test")
	fmt.Fprintln(w, "URL.Path", r.URL.Path)
	fmt.Fprintln(w, "URL parameters", r.Form)
	fmt.Fprintln(w, "Form prmt1", r.Form["prmt1"])

	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}

}

func manageLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		username := r.FormValue("username")
		fmt.Println("username:", username)
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, strings.Join([]string{"Hello ", username, "!"}, ""))
	}
}
