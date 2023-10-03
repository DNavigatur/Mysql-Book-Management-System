package routes

import (
	"github.com/DNavigatur/bookstore-API/pkg/controllers"
	"github.com/gorilla/mux"
)

// This is defining the routes for each endpoints
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("POST")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
