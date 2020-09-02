package model

import (
	dbConn "rfw-bend/config/database"
)

func userLogin(username string, password string) int {

	conn := dbConn.ArangoDBConnect()
	conn.NewQuery("FOR d IN test RETURN d")
	return 0
}
