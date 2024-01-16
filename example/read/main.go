package main

import (
	"context"
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

	getContentParams := sheetdb.GetContentParams{}

  ctx := context.TODO()
	contents, err := sheetDB.Read.GetContentWithContext(ctx, &getContentParams)

	if err != nil {
		log.Fatal(err)
	}

	for _, content := range *contents {
		log.Println(content)
	}
}
