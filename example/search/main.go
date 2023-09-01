package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rslhdyt/sheetdb-go"
	"github.com/rslhdyt/sheetdb-go/client"
)

func main() {
	godotenvErr := godotenv.Load()

	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}

	sheetDBUsername := os.Getenv("SHEETDB_USERNAME")
	sheetDBPassword := os.Getenv("SHEETDB_PASSWORD")
	sheetDBDocumentId := os.Getenv("SHEETDB_DOCUMENT_ID")

	sheetDB := client.New(sheetDBUsername, sheetDBPassword, sheetDBDocumentId)

	searchParams := sheetdb.SearchParams{
		"Name": "Rumi",
	}

	content, err := sheetDB.Search.Find(&searchParams)

	if err != nil {
		log.Fatal(err)
	}

  log.Println(content)
}
