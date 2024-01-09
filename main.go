package main

import (
	"ZWorld/internal/web"
)

func main() {
	server := web.RegisterRoutes()

	server.Run(":8080")
}
