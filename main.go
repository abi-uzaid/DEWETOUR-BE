package main

import (
	"dewetour/database"
	"dewetour/pkg/mysql"
	"dewetour/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.DatabaseInit()
	database.RunMigration()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.RouteInit(r.Group("/api/v1"))

	fmt.Println("Server started")
	http.ListenAndServe("localhost:5000", r)

}
