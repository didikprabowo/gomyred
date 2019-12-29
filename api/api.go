package api

import (
	"context"
	"fmt"
	"github.com/gomyred/config"
	"github.com/gomyred/utils"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	cfg config.Config
	lg  utils.Logger
)

// RunningAPI
func RunningAPI() {
	cfg = config.GetConfig()

	m := http.NewServeMux()

	s := http.Server{
		Addr:         fmt.Sprintf(":%d", +cfg.Server.Port),
		Handler:      m,
		ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
		WriteTimeout: cfg.Server.WriteTimeout * time.Second,
	}

	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Println("Server running...")
		for {
			select {
			case <-ch:
				s.Shutdown(context.Background())
			case <-ctx.Done():
				return
			}
		}
	}()

	m.HandleFunc("/news", GetArticle)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		lg.Info = "Server"
		lg.Error = err
		lg.NewLogError()
		fmt.Println(err)

	}
	fmt.Println("Stop server...")
}
