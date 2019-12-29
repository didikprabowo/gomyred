package api

import (
	"encoding/json"
	"fmt"
	"github.com/gomyred/config"
	"github.com/gomyred/src/article/model"
	"github.com/gomyred/src/article/repository"
	"github.com/gomyred/utils"
	"net/http"
	"strconv"
	"sync"
)

var (
	articles []model.Article
	dconf    config.Config
)

type newsFetchContainer struct {
	Stack []model.Article `json:"Stack"`
}

func init() {
	dconf = config.GetConfig()
}

// GetArticle
func GetArticle(w http.ResponseWriter, r *http.Request) {
	var start, end, count int
	resultOut := make([]model.Article, 0)
	count = 0

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	if page == 1 {
		start = 0
		end = 10
	} else {
		start = (page - 1) * 10
		end = 10
	}

	checkChace, _ := repository.GetCache(page)
	if len(checkChace) > 0 {
		json.Unmarshal([]byte(checkChace), &articles)
		resultOut = articles
		count = len(checkChace)
		fmt.Println("Data From Redis...")
	} else {
		fmt.Println("Data From  MySQL...")
		jobs := make(chan model.Article, 1)
		var wg sync.WaitGroup
		var m = new(sync.Mutex)

		go func() {
			wg.Add(1)
			repository.GetAllArticle(jobs, &wg, start, end)
		}()

		for result := range jobs {
			m.Lock()
			count++
			resultOut = append(resultOut, result)
			m.Unlock()
		}
		toJson, err := json.Marshal(resultOut)
		if err != nil {
			fmt.Println(err)
		}
		repository.SetCache(page, toJson, dconf.Server.TTL)

		wg.Wait()
	}

	out := map[string]interface{}{
		"results": resultOut,
		"meta": map[string]interface{}{
			"per_page":     10,
			"current_page": page,
			"count_page":   count,
		},
	}

	utils.RespondJSON(w, http.StatusOK, out)

}
