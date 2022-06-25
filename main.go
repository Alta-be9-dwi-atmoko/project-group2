package main

import (
	"project/group2/config"
	"project/group2/factory"
	"project/group2/migration"
	"project/group2/routes"
)

func main() {
	//initiate db connection
	dbConn := config.InitDB()

	// run auto migrate from gorm
	migration.InitMigrate(dbConn)

	// initiate factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":80"))
}
