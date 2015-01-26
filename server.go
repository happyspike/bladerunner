package bladerunner

import (
	"log"
	"net/http"
	"os"
)

func Run(handler http.Handler) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := ":" + port
	log.Println("Starting server on " + addr)
	http.ListenAndServe(addr, handler)
}
