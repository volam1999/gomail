package main

import (
	"net/http"

	"github.com/volam1999/gomail/internal/app/api"
	"github.com/volam1999/gomail/internal/pkg/config/envconfig"
	"github.com/volam1999/gomail/internal/pkg/log"
)

func main() {
	if envconfig.SetEnvFromFile("configs/config.env") != nil {
		log.Fatal("Error loading .env file")
		return
	}

	router := api.NewRouter()
	http.ListenAndServe(":8080", router)
}
