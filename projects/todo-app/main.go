package main
import (
	"net/http"
	"html/template"
	"fmt"
)

var temp *template.Template



type Todo struct {
	Item 		string
	Done  		bool
}

type TodoList struct {
	Title 		string
	TodoList 	[]Todo

}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	data := TodoList{
		Title: "TODO List",
		TodoList: []Todo{
			{Item: "Write TODO app", Done: true},
			{Item: "Build a web server using GO-Lang", Done: true},
			{Item: "Deploy on Kubernetes", Done: false},
		},
	}

	temp.Execute(w, data)
}


func staticHandler(path string){
	fs := http.FileServer(http.Dir(fmt.Sprintf("./%s", path)))
	http.Handle("/static/", http.StripPrefix(fmt.Sprintf("/%s/", path), fs))
}

func servForEver(port int) {
	fmt.Println("Server started on port:", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}

func main() {
	temp = template.Must(template.ParseFiles("templates/index.html"))
	staticHandler("static")
	http.HandleFunc("/", mainPage)
	servForEver(8080)
}