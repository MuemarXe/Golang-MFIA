package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/MuemarXe/go-bookstore/pkg/routes"

)
func main(){
	r := mux.NewRouter()
	routes.RegistersBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost9010",r))

}