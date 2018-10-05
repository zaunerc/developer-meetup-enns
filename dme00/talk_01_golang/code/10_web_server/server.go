package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	cwd, _ := os.Getwd()

    // ServeMux is an HTTP request multiplexer.
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(cwd)))

	addr := "localhost:3000"
	log.Printf("Server will be listening on port %v", addr)
	log.Printf("Server will be serving dir %v", cwd)

	/*
	 * http.ListenAndServe opens the server port, and blocks
	 * forever waiting for clients. If it fails to open the
	 * port, the log.Fatal call will report the problem and
	 * exit the program.
	 */
	log.Fatal(http.ListenAndServe(addr, mux))
}
