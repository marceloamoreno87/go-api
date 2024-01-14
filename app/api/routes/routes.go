package routes

import (
	"net/http"
	"os"

	_ "github.com/marceloamoreno/izimoney/docs"
	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetUserRoutes(m *http.ServeMux) {
	m.HandleFunc("/users", handler.GetUsers)
	m.HandleFunc("/user", handler.GetUser)
	m.HandleFunc("/user/create", handler.CreateUser)
	m.HandleFunc("/user/update", handler.UpdateUser)
	m.HandleFunc("/user/delete", handler.DeleteUser)
}

func GetSwaggerRoutes(m *http.ServeMux) {
	m.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+os.Getenv("PORT")+"/swagger/doc.json"), //The url pointing to API definition
	))
}
