package database

import (
	"context"

	ArangoDriver "github.com/arangodb/go-driver"
	ArangoHttp "github.com/arangodb/go-driver/http"
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

// SettingDB - Set config
type SettingDB struct {
	dbConn   ArangoDriver.Connection
	dbClient ArangoDriver.ClientUsers
	dbSelect ArangoDriver.Database
	dbStatus string
}

// CursorQuery - Set config Curcor
type CursorQuery struct {
	dbCursor    ArangoDriver.Cursor
	QueryStatus string
}

// ArangoDBConnect - Open connect to the database
func ArangoDBConnect() *SettingDB {

	connDB, connErr := ArangoHttp.NewConnection(ArangoHttp.ConnectionConfig{
		Endpoints: []string{DBurl}},
	)

	if connErr != nil {
		// Handle error
	}

	clientDB, clientErr := ArangoDriver.NewClient(ArangoDriver.ClientConfig{
		Connection:     connDB,
		Authentication: ArangoDriver.BasicAuthentication(DBuser, DBpass),
	})

	if clientErr != nil {
		// Handle error
	}

	selectDB, errSelect := clientDB.Database(nil, DBselect)

	if errSelect != nil {
		// Handle error
	}

	return &SettingDB{
		dbConn:   connDB,
		dbClient: clientDB,
		dbSelect: selectDB,
		dbStatus: "OK",
	}
}

//NewQuery - Query Database
func (db *SettingDB) NewQuery(query string) *CursorQuery {

	ctx := ArangoDriver.WithQueryCount(context.Background())
	dbQuery, errQuery := db.dbSelect.Query(ctx, query, nil)

	if errQuery != nil {

	}

	return &CursorQuery{
		dbCursor:    dbQuery,
		QueryStatus: "OK",
	}
}

/*

type ServiceClientDB struct {
	dbClient ArangoDriver.ClientUsers
	dbStatus string
}

type QueryClientDB struct {
	dbQuery  ArangoDriver.Cursor
	dbStatus string
}


func NewService(dbc ArangoDriver.Client) *ServiceClientDB {

	dbc.Database(nil, DBselect)

	return &ServiceClientDB{
		dbClient: dbc,
		dbStatus: "OK",
	}
}

func NewQuery(dbq ArangoDriver.Database, query string) *QueryClientDB {

	var statusQuery string = "QUERY-OK"
	ctx := ArangoDriver.WithQueryCount(context.Background())
	dbcursor, errcursor := dbq.Query(ctx, query, nil)

	if errcursor != nil {
		statusQuery = "QUERY-ERROR"
	}

	return &QueryClientDB{
		dbQuery:  dbcursor,
		dbStatus: statusQuery,
	}

}
*/
