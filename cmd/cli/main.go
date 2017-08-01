package main

import (
	"fmt"
	"log"

	"github.com/dave-malone/aws-sandbox/pkg/config"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.NewConfig()
	fmt.Printf("config: %v\n", config)

	// db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
