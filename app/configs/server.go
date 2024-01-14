package configs

import (
	"net/http"
	"os"

	"github.com/marceloamoreno/izimoney/internal/domain/user/routes"
)

func StartServer() {
	mux := http.NewServeMux()
	loadRoutes(mux)
	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}

func loadRoutes(m *http.ServeMux) {
	routes.GetUserRoutes(m)
}
