package main

import (
	"fmt"
	"hafiztri123/be-sharing-vision/internal/database"

	"github.com/joho/godotenv"
)



func main() {
	//Relative to where you run the program
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()
	defer db.Close()


	fmt.Println("Database connected")
}