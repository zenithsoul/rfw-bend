package main

import (
	"fmt"
	testdb "rfw-bend/config/database"
)

// "net/http"

// "github.com/labstack/echo"

func main() {

	//

	/*
		e := echo.New()
		e.GET("/", func(c echo.Context) error {

			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Fatal(e.Start(":8080"))
	*/

	fmt.Println(testdb.ArgoGetInfoDB())
}
