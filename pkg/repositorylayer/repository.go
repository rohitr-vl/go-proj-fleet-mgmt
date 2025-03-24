package repositorylayer

import (
	db_sqlc "fleetmgmt/web/internal/db/sqlc"
	"log"
)

func runDb() {
	if err := db_sqlc.InitiateDbConnection(); err != nil {
		log.Fatal(err)
	}
}