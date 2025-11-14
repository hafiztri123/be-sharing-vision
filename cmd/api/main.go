package main

import (
	"fmt"
	"hafiztri123/be-sharing-vision/internal/articles"
	"hafiztri123/be-sharing-vision/internal/database"
	"hafiztri123/be-sharing-vision/internal/router"
	"hafiztri123/be-sharing-vision/internal/utils"
	"net/http"

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
	articleRepository := articles.NewRepository(db)
	articleHandler := articles.NewHandler(articleRepository)

	mux := router.NewRouter(articleHandler)




	fmt.Println("Database connected")

	http.ListenAndServe(fmt.Sprintf(":%s", utils.GetMandatoryEnv("APP_PORT")), mux)


}