package main

// "net/http"

// "github.com/labstack/echo"

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	dbConn "rfw-bend/config/database"

	"github.com/arangodb/go-driver"
)

// MyDocument -
type MyDocument struct {
	ID         string `json:"_id"`
	Key        string `json:"_key"`
	Rev        string `json:"_rev"`
	ValuePrice int64  `json:"value"`
}

func main() {

	start := time.Now()

	conn := dbConn.ArangoDBConnect()
	getCursor, _, _ := conn.NewQuery("FOR d IN test RETURN d")
	defer getCursor.Close()
	getCount := getCursor.Count()

	strtxt := strconv.FormatInt(getCount, 10)
	fmt.Println("total = ", strtxt)

	result := []MyDocument{}

	for {
		product := MyDocument{}
		_, err := getCursor.ReadDocument(context.Background(), &product)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}

		result = append(result, product)
	}

	println(result[99].ValuePrice)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	println(len(result))

}
