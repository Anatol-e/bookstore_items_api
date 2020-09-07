package app

import (
	"github.com/Anatol-e/bookstore_items_api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemControllers.Create).Methods(http.MethodPost)
}
