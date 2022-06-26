package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/swaggo/http-swagger"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "gokiosk/api"
	"log"
	"net/http"
)

// @title Warehouse Management API
// @version 1.0
// @description This is a simple Warehouse management project for learning Go language and PostgreSQL database.
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api
func main() {
	readDotEnv()
	port := viper.Get("APP_PORT").(string)

	boil.DebugMode = true
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"), //The url pointing to API definition
	))

	welcomeApp(port)
	log.Fatalln(http.ListenAndServe(":"+port, r))
}

func readDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func welcomeApp(port string) {
	fmt.Println(" ██████╗  ██████╗ ██████╗ ██╗  ██╗███████╗██████╗ ")
	fmt.Println("██╔════╝ ██╔═══██╗██╔══██╗██║  ██║██╔════╝██╔══██╗")
	fmt.Println("██║  ███╗██║   ██║██████╔╝███████║█████╗  ██████╔╝")
	fmt.Println("██║   ██║██║   ██║██╔═══╝ ██╔══██║██╔══╝  ██╔══██╗")
	fmt.Println("╚██████╔╝╚██████╔╝██║     ██║  ██║███████╗██║  ██║")
	fmt.Println(" ╚═════╝  ╚═════╝ ╚═╝     ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝")
	log.Printf("Server is running on port %s\n", port)
}
