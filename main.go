package main

import (
	db "github.com/ArisFrsy/go-backend/database"
	fiberMain "github.com/ArisFrsy/go-backend/route"
)

func main() {
	db.InitDB()
	db.AutoMigrateModels()
	fiberMain.FiberHandler()
}
