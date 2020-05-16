package app

import (
	"github.com/alessandroarosio/bookstore_users-API/controllers/ping"
	"github.com/alessandroarosio/bookstore_users-API/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
