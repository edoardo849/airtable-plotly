package main

import (
	"fmt"
	"os"

	airtable "github.com/fabioberger/airtable-go"
)

type idea struct {
	AirtableID string
	Fields     struct {
		Name          string  `json:"Name"`
		Category      string  `json:"Category"`
		Impact        float32 `json:"Impact score"`
		Achievability float32 `json:"Achievability score"`
	}
}

// Impact, achievability, category

func main() {
	airtableAPIKey := os.Getenv("AIRTABLE_API_KEY")
	baseID := "app9nxg1VNqQXUOZA"

	client, err := airtable.New(airtableAPIKey, baseID)
	if err != nil {
		panic(err)
	}

	listParams := airtable.ListParameters{
		Fields:          []string{"Name", "Category", "Impact score", "Achievability score"},
		FilterByFormula: "AND({Revenue-generating proposition?} = \"Yes\", {B2B offering?} = \"Yes\", {Non-headcount scalable?} = \"Yes\", {Better with Deloitte?} = \"Yes\", {Right to play?} = \"Yes\")",
		MaxRecords:      1000,
		View:            "Main View",
	}
	ideas := []idea{}
	if err := client.ListRecords("Main database", &ideas, listParams); err != nil {
		panic(err)
	}

	fmt.Println(ideas)

}
