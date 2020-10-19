package main

import (
	"encoding/json"
	"fmt"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type TodoItem struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
}

var todoItems []TodoItem = []TodoItem{
	TodoItem{Id: 1, Todo: "clean the house"},
	TodoItem{Id: 2, Todo: "wash car"},
	TodoItem{Id: 3, Todo: "go shoppping"},
	TodoItem{Id: 4, Todo: "cook dinner"},
}

func getTodoItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonBytes, _ := json.Marshal(todoItems)
	fmt.Fprintf(w, string(jsonBytes)+"\n")
}

func saidDaniel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello my name is joe\n")
}

func saidJamie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "keep working daniel.\n")
}

func initHandlers(r *mux.Router) {
	r.HandleFunc("/todo-item", getTodoItems).Methods("GET")
	r.HandleFunc("/daniel", saidDaniel).Methods("GET")
	r.HandleFunc("/jamie", saidJamie).Methods("GET")
}

func main() {
	router := mux.NewRouter()
	initHandlers(router)
	log.Println("Listening on :8000")
	log.Println(http.ListenAndServe(":8000", ghandlers.CORS(
		ghandlers.AllowedHeaders([]string{
			"Access-Control-Allow-Origin",
			"X-Requested-With",
			"Content-Type",
			"Authorization",
		}),
		ghandlers.AllowedMethods([]string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"HEAD",
			"OPTIONS",
		}),
		ghandlers.AllowedOrigins([]string{
			os.Getenv("ORIGIN_ALLOWED"),
			"*",
		},
		))(router)))
}
