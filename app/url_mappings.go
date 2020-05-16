package app

import (
	"github.com/alessandroarosio/bookstore_users-API/controllers/ping"
	"github.com/alessandroarosio/bookstore_users-API/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
