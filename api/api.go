package api

import (
	"fmt"
	"github.com/gomyred/config"
	"net/http"
)

var cfg config.Config

// RunningAPI
func RunningAPI() {
	cfg = config.GetConfig()
	urlServer := fmt.Sprintf(":%d", +cfg.Server.Port)

	http.HandleFunc("/news", GetArticle)
	http.ListenAndServe(urlServer, nil)
}
