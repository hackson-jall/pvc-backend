package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	port := 3000
	fmt.Println("Starting Server on port", port)

	http.HandleFunc("/get_mod_3", func(w http.ResponseWriter, r *http.Request) {
		EnableCors(&w)
		r.ParseForm()
		value := r.Form["value"][0]
		int_value, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		int_value = int_value % 3
		w.Write([]byte("Value: " + fmt.Sprint(int_value)))
	})
	http.ListenAndServe(":3000", nil)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") //TODO beef up security
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
