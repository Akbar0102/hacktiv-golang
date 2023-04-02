package main

import (
	"assignment2/database"
	"assignment2/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()

	routers.StartServer().Run(PORT)
}
