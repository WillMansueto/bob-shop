package main

import(
	"fmt"
	"log"
	"os"
	"net/http"
	"go-webapp/routes"
	"go-webapp/utils"
	"go-webapp/models"
)

func main(){
	models.TestConnection()

	port := os.Getenv("PORT")
	if port =="" {
		fmt.Println("Not Port")
		os.Exit(1)
	}
	fmt.Printf("listening Port %s\n", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
