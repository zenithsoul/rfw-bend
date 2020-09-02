package database

import (
	"context"

	ArangoDriver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

const (
	// DBurl -
	DBurl string = "http://localhost:8529"
	// DBuser -
	DBuser string = "root"
	// DBpass -
	DBpass string = "1234"
	// DBselect -
	DBselect string = "firewall"
)

type ServiceClientDB struct {
	dbClient ArangoDriver.ClientUsers
}

type QueryClientDB struct {
	dbQuery ArangoDriver.Cursor
}

func NewService(dbc ArangoDriver.Client) *ServiceClientDB {

	dbc.Database(nil, DBselect)

	return &ServiceClientDB{
		dbClient: dbc,
	}
}

func NewQuery(dbq ArangoDriver.Database, query string) *QueryClientDB {

	ctx := ArangoDriver.WithQueryCount(context.Background())
	dbq.Query(ctx, query, nil)

}

// ArangoDBConnect - Open connect to the database
func ArangoDBConnect() (ArangoDriver.Client, error) {

	conndb, connerr := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBurl}},
	)

	if connerr != nil {
		// Handle error
	}

	clientdb, clienterr := ArangoDriver.NewClient(ArangoDriver.ClientConfig{
		Connection:     conndb,
		Authentication: ArangoDriver.BasicAuthentication(DBuser, DBpass),
	})

	if clienterr != nil {
		// Handle error
	}

	return clientdb, nil
}

//DBconnect - database that we connected
/*
func DBconnect() {

	ctx := ArangoDriver.WithQueryCount(context.Background())

	conndb, connerr := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBurl}},
	)

	if connerr != nil {
		// Handle error
	}

	clientdb, clienterr := ArangoDriver.NewClient(ArangoDriver.ClientConfig{
		Connection:     conndb,
		Authentication: ArangoDriver.BasicAuthentication(DBuser, DBpass),
	})

	if clienterr != nil {
		// Handle error

	}

	db, dberr := clientdb.Database(nil, DBselect)

	return db
}
*/
