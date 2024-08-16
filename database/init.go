package database

import (
	"app/database/psql"
)

func init(){
	psql.Connect()
}

func SetupDB() {}