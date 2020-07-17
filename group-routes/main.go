package main

import (
	"github.com/gin-gonic/examples/group-routes/routes"
)

func main() {
	// Our server will live in the routes package
	routes.Run()
}
