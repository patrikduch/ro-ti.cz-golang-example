package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	content := `Test GoLang BE project`

	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
