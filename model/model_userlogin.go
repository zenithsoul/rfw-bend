package model

import (
	"rfw-bend/config/database"
)

func userLogin(username string, password string) int {

	conn, err := database.ArangoDBConnect()

	if err != nil {

	}

	database.NewService(conn)

	return 0
}
