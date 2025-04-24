package main

import (
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/aoemedia-server/adapter/driving/restful/route"
)

func main() {
	db.InitDB()
	route.InitEngine()
}
