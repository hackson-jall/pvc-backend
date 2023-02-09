package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ingredient struct {
	Name     string   `json:"name"`
	Quantity Quantity `json:"quantity"`
}
type Quantity struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
}

func ParseRecipe(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	url := r.Form.Get("url")

	ingredients := []Ingredient{
		{
			Name: "Apple",
			Quantity: Quantity{
				Amount: "1",
				Unit:   "g",
			},
		},
		{
			Name: "Orange",
			Quantity: Quantity{
				Amount: "4",
				Unit:   "items",
			},
		},
		{
			Name: "Banana",
			Quantity: Quantity{
				Amount: "3",
				Unit:   "bunches",
			},
		},
	}
	outputJson, err := json.Marshal(ingredients)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", outputJson)
}
