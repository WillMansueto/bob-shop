package main

import(
	"fmt"
	"log"
	"net/http"
	"go-webapp/routes"
	"go-webapp/utils"
	"go-webapp/models"
)

const PORT = ":8080"

func main(){
	models.TestConnection()
	fmt.Printf("listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
