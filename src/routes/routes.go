package routes

import (
	"github.com/gorilla/mux"
	"github.com/ojaoferreira/go-rest-api/src/controllers"
	"github.com/ojaoferreira/go-rest-api/src/middleware"
)

func HandleRequest() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.ListaTodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ListaPersonalidadePorId).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaUmaPersonalidade).Methods("Put")
	return r
}
