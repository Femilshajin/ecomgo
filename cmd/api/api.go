package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/femilshajin/ecomgo/cmd/service/product"
	"github.com/femilshajin/ecomgo/cmd/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)

	v1 := http.NewServeMux()
	v1.Handle("/", userHandler.RegisterRoutes())
	v1.Handle("/products", productHandler.RegisterRoutes())

	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Server started at the port: " + s.addr)

	return server.ListenAndServe()
}
