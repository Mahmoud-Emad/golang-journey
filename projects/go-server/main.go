package main 

import(
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name:= r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandeler(w http.ResponseWriter, r *http.Request){
	// hello string = "/hello"
	// hello_slash string = "/hello"
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found Page", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not allowed.", http.StatusNotFound)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandeler)

	fmt.Printf("Starting server under port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}