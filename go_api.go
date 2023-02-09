package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func main() {
	recipeParserFunc := http.HandlerFunc(ParseRecipe)
	compressedRecipeParserFunc := gziphandler.GzipHandler(recipeParserFunc)
	http.Handle("/parse_recipe", compressedRecipeParserFunc)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") //TODO beef up security
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
