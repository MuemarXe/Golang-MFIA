package routes
import(
	//"fmt"
	"github.com/gorilla/mux"

	"github.com/MuemarXe/go-bookstore/pkg/controllers"
)

var RegistersBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	
}