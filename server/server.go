package server

import (
	"log"
	"net/http"
	"os"

	"github.com/rtravitz/culture_knights/router"
)

func StartServer() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
