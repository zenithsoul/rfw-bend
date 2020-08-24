package database

import (
	"context"
	"log"
	"strconv"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type server struct {
	db *sql.DB,
	
}


// ArgoGetInfoDB - Get Test
func ArgoGetInfoDB() string {

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
		return "Connection TCP Error"
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "1234"),
	})

	if err != nil {
		// Handle error
		return "Connection Driver Error"
	}

	db, err := c.Database(context.Background(), "firewall")

	if err != nil {
		// Handle error
		return "Connection DB Error"
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
	log.Printf("Binomial took %s", elapsed)

	return strtxt
}
