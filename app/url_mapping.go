package app

import (
	"github.com/S4ND1X/BookStoreUsersAPI/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
