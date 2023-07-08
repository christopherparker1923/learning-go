package main

import(
	"log"
	"net/http"
	"github.com.gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/christopherparker1923/learning-go/bookstore-apis/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010",))
}