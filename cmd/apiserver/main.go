package main

import (
	"flag"
	"log"

	"erzi_new/internal/apiserver"
	producthalder "erzi_new/internal/handler/product"
	productrepo "erzi_new/internal/repository/product"
	productsrv "erzi_new/internal/service/product"
	"erzi_new/internal/store"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	db := store.New()
	err = db.Open(config.Store.DatabaseURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	productRepo := productrepo.NewRepository(db)
	productSrv := productsrv.NewService(productRepo)
	productHandler := producthalder.NewHandler(productSrv)
	s := apiserver.New(config)
	s.ConfigureRouter(productHandler)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
