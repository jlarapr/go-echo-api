package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println("Welcome golang")

	server := echo.New()
	server.GET(path.Join("/"), Version)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	err = server.Start(address)
	if err != nil {
		panic(err)
	}
}

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 1})
}
