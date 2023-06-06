package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"ngmi_server/internal/config"
	"ngmi_server/internal/health_check"
	"ngmi_server/internal/router"
	database "ngmi_server/pkg/db"
	"ngmi_server/pkg/log"
	"os"
)

const Version = "0.0.1"

var configFlag = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()

	logger := log.New().With(nil, "version", Version)
	if cfg, err := config.Load(*configFlag); err != nil {
		logger.Error(err)
		os.Exit(-1)
	} else {
		if db, err := sqlx.Open("mysql", cfg.DSN); err != nil {
			logger.Error(err)
			os.Exit(-1)
		} else {
			defer func() {
				if err := db.Close(); err != nil {
					logger.Error(err)
				}
			}()
			port := fmt.Sprintf(":%v", cfg.ServerPort)
			dtb := database.New(db)
			if err := dtb.Migrate("C:\\Users\\Anas Youssef\\Documents\\projects\\ngmi_server\\migrations"); err != nil {
				logger.Error(err)
				os.Exit(-1)
			}
			r := router.BuildHandler(dtb, logger)
			health_check.RegisterHandlers(r, Version)
			hs := &http.Server{Addr: port, Handler: r}
			logger.Infof("Server running at %v", port)
			if err = hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error(err)
				os.Exit(-1)
			}
		}
	}
}
