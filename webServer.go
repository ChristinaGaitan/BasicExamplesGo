//Web server

package main

import (
//	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // cada petición manda llamar un handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// Redireccionar a un página
	//http.Redirect(w,r, "http://www.google.com", 301)

	// Otra manera de redireccionar
	http.Redirect(w,r, "http://www.google.com", http.StatusMovedPermanently)
}