package main

import (
	"context"
	"log"
	"strconv"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
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

	// fmt.Println(testdb.ArgoGetInfoDB())

	start := time.Now()

	//r := new(big.Int)
	//fmt.Println(r.Binomial(1000, 10))

	ctxx := context.Background()
	ctx := driver.WithQueryCount(ctxx)

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})

	if err != nil {
		// Handle error

	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "1234"),
	})

	if err != nil {
		// Handle error

	}

	db, err := c.Database(nil, "firewall")

	if err != nil {
		// Handle error
	}

	query := "FOR d IN test RETURN d"
	cursor, err := db.Query(ctx, query, nil)

	defer cursor.Close()

	if err != nil {
		// handle error
	}

	scorunt := cursor.Count()
	strtxt := strconv.FormatInt(scorunt, 10)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s : total %s", elapsed, strtxt)
}
