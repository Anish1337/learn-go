package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlePage)

	// Read the PORT from the environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	addr := ":" + port

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>`
	w.Write([]byte(page))
}

