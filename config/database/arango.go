package database

import (
	"context"

	ArangoDriver "github.com/arangodb/go-driver"
	ArangoHttp "github.com/arangodb/go-driver/http"
)

const (
	// DBurl -
	DBurl string = "http://192.168.10.188:8529"
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
	dbError  uint8
}

// ArangoDBConnect - Open connect to the database
func ArangoDBConnect() *SettingDB {

	var dbStatus string = "Success"
	var dbError uint8 = 0

	connDB, connErr := ArangoHttp.NewConnection(ArangoHttp.ConnectionConfig{
		Endpoints: []string{DBurl}},
	)

	if connErr != nil {
		// Handle error
		dbStatus = "Error: TCP Connect"
		dbError = 1
	}

	clientDB, clientErr := ArangoDriver.NewClient(ArangoDriver.ClientConfig{
		Connection:     connDB,
		Authentication: ArangoDriver.BasicAuthentication(DBuser, DBpass),
	})

	if clientErr != nil {
		// Handle error
		dbStatus = "Error: Client Connection"
		dbError = 2
	}

	selectDB, selectErr := clientDB.Database(nil, DBselect)

	if selectErr != nil {
		// Handle error
		dbStatus = "Error: Database Connection"
		dbError = 3
	}

	return &SettingDB{
		dbConn:   connDB,
		dbClient: clientDB,
		dbSelect: selectDB,
		dbStatus: dbStatus,
		dbError:  dbError,
	}
}

//NewQuery - Query Database and return a data to Cursor (Interface) , DB status and Error code
func NewQuery(query string) (ArangoDriver.Cursor, string, uint8) {

	var db SettingDB = *ArangoDBConnect()
	ctx := ArangoDriver.WithQueryCount(context.Background())
	dbQuery, errQuery := db.dbSelect.Query(ctx, query, nil)

	if errQuery != nil {
		db.dbStatus = "Error: Query"
		db.dbError = 4
	}

	return dbQuery, db.dbStatus, db.dbError
}

/*
func (db *SettingDB) NewQuery(query string) (ArangoDriver.Cursor, string, uint8) {

	ctx := ArangoDriver.WithQueryCount(context.Background())
	dbQuery, errQuery := db.dbSelect.Query(ctx, query, nil)

	if errQuery != nil {
		db.dbStatus = "Error: Query"
		db.dbError = 4
	}

	return dbQuery, db.dbStatus, db.dbError
}
*/
