package routes

import (
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
)

func GetUserRoutes(m *http.ServeMux) {
	m.HandleFunc("/users", handler.GetUsers)
	m.HandleFunc("/user", handler.GetUser)
	m.HandleFunc("/user/create", handler.CreateUser)
	m.HandleFunc("/user/update", handler.UpdateUser)
	m.HandleFunc("/user/delete", handler.DeleteUser)
}
