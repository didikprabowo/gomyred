package main_test

import (
	"fmt"
	"github.com/gomyred/api"
	"github.com/gomyred/config"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var (
	cfg config.Config
)

func tesHTTP(t *testing.T) {
	cfg = config.GetConfig()

	url := fmt.Sprintf("http://127.0.0.1:%d/news", cfg.Server.Port)

	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	api.GetArticle(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	t.Logf(string(body))
}
