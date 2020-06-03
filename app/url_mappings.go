package app

import (
	"github.com/PiotrJanczewski/bookstore_users_api/controllers/ping"
	"github.com/PiotrJanczewski/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// USER ROUTES
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
