package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func responseMessage() string {
	feeling := os.Getenv("FEELING")
	if feeling == "" {
		feeling = "$FEELING"
	}

	awesomness := os.Getenv("AWESOMNESS")
	if awesomness == "" {
		awesomness = "$AWESOMNESS"
	}

	return fmt.Sprintf("Howdy! I feeeeel %s today. I will tell you my secret of awesomeness: \"%s\".", feeling, awesomness)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, responseMessage())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server running on http://%s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
