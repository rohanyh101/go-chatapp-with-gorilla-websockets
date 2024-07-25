package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func initFrontEnd(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("./public"))

	fs.ServeHTTP(w, r)
}

func main() {
	port := 8080
	manager := NewManager()
	mux := http.NewServeMux()

	mux.HandleFunc("/", initFrontEnd)
	mux.HandleFunc("/ws", manager.serveWS)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		Handler:           mux,
	}

	log.Printf("server running on port: %d", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
