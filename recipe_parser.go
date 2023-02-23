package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	re "regexp"
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
	EnableCors(&w)
	err := r.ParseForm()
	recipeURL := r.Form.Get("url")
	requestURL := "https://recipe-parser.azurewebsites.net/api/parse?url=" + recipeURL
	resp, err := http.Get(requestURL)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}
	ingredientsRaw := result["recipeIngredient"].([]interface{})
	ingredients := []Ingredient{}
	for _, ingredientStr := range ingredientsRaw {
		parts := re.MustCompile("\\s\\s+").Split(ingredientStr.(string), -1)
		ingredient := Ingredient{
			Name:     parts[2],
			Quantity: Quantity{Amount: parts[0], Unit: parts[1]},
		}
		ingredients = append(ingredients, ingredient)

	}
	outputJson, err := json.Marshal(ingredients)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", outputJson)
}
