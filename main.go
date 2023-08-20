package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"                                     ❶

	mux := http.NewServeMux()

	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {

			enc := json.NewEncoder(w)
			w.Header().
				Set("Content-Type",
					"application/json; charset=utf-8")
			resp := Resp{
				Language:    "English",
				Translation: "Hello",
			}
			if err := enc.Encode(resp); err != nil {
				panic(fmt.Sprintf("Error %s", err.Error()))
			}
		})

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}


