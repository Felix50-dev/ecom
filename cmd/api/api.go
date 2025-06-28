package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Felix50-dev/ecom/service/cart"
	"github.com/Felix50-dev/ecom/service/order"
	"github.com/Felix50-dev/ecom/service/product"
	"github.com/Felix50-dev/ecom/service/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter().StrictSlash(true)
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, userStore, productStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
