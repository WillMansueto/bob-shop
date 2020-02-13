package main

import(
	"fmt"
	"log"
	"os"
	"net/http"
	"bob-shop/routes"
	"bob-shop/utils"
	"bob-shop/models"
)

func main(){
	models.TestConnection()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("listening Port %s\n", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
