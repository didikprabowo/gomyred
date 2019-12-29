package config

import (
	"fmt"
	"github.com/JeremyLoy/config"
	"sync"
	"time"
)

type (
	Server struct {
		Port int `config:"SERVER_PORT"`
		TTL  int `config:"CACHE_TTL"`
	}

	MySQL struct {
		Host     string `config:"MYSQL_HOST"`
		Database string `config:"MYSQL_DATABASE"`
		Username string `config:"MYSQL_USERNAME"`
		Password string `config:"MYSQL_PASSWORD"`
		Port     int    `config:"MYSQL_PORT"`
	}

	Redis struct {
		Host     string `config:"REDIS_HOST"`
		Database int    `config:"REDIS_DATABASE"`
		Password string `config:"REDIS_PASSWORD"`
		Port     int    `config:"REDIS_PORT"`
	}

	Config struct {
		Server Server
		MySQL  MySQL
		Redis  Redis
	}
)

var (
	srv   Server
	mysql MySQL
	redis Redis
	cfg   Config
)

// DSN
func (m MySQL) DSN() string {
	url := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database)
	return url
}

// GetConfig
func GetConfig() Config {

	c := make(chan Config, 2)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		config.From(".env").To(&cfg.Server)
		c <- cfg
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		config.From(".env").To(&cfg.MySQL)
		c <- cfg
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		config.From(".env").To(&cfg.Redis)
		c <- cfg
	}()

	timeout := time.Second

	select {
	case <-c:
		fmt.Println("")
	case <-time.After(timeout):
		fmt.Println("time out om")
	}

	wg.Wait()

	return cfg

}
