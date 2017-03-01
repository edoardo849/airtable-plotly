package main

import (
	"net/http"

	airtable "github.com/fabioberger/airtable-go"
)

type fields struct {
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Impact        float32 `json:"Impact score"`
	Achievability float32 `json:"Achievability score"`
}

type fieldsOut struct {
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Impact        float32 `json:"impact"`
	Achievability float32 `json:"achievability"`
}

type idea struct {
	AirtableID string `json:"ID"`
	Fields     fields `json:"fields"`
}

func (s *Server) handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handlePollsGet(w, r)
		return
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		respond(w, r, http.StatusOK, nil)
		return
	}
	// not found
	respondHTTPErr(w, r, http.StatusNotFound)
}

func (s *Server) handlePollsGet(w http.ResponseWriter, r *http.Request) {
	//if err := q.All(&result); err != nil {
	//	respondErr(w, r, http.StatusInternalServerError, err)
	//	return
	//}

	baseID := "app9nxg1VNqQXUOZA"

	client, err := airtable.New(s.APIKey, baseID)
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}

	listParams := airtable.ListParameters{
		Fields:          []string{"Name", "Category", "Impact score", "Achievability score"},
		FilterByFormula: "AND({Revenue-generating proposition?} = \"Yes\", {B2B offering?} = \"Yes\", {Non-headcount scalable?} = \"Yes\", {Better with Deloitte?} = \"Yes\", {Right to play?} = \"Yes\")",
		MaxRecords:      1000,
		View:            "Main View",
	}
	ideas := []idea{}

	if err := client.ListRecords("Main database", &ideas, listParams); err != nil {

		respondErr(w, r, http.StatusInternalServerError, err)
		return

	}

	var res []fieldsOut
	for _, in := range ideas {
		out := fieldsOut(in.Fields)
		res = append(res, out)
	}

	respond(w, r, http.StatusOK, res)
}
