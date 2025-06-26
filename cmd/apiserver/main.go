package main

import (
	"erzi_new/internal/apiserver"
	cartHalder "erzi_new/internal/handler/cart"
	producthalder "erzi_new/internal/handler/product"
	userhalder "erzi_new/internal/handler/user"
	cartrepo "erzi_new/internal/repository/cart"
	productrepo "erzi_new/internal/repository/product"
	userrepo "erzi_new/internal/repository/user"
	cartsrv "erzi_new/internal/service/cart"
	productsrv "erzi_new/internal/service/product"
	usersrv "erzi_new/internal/service/user"
	"erzi_new/internal/store"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
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
	userRepo := userrepo.NewRepository(db)
	userSrv := usersrv.NewService(userRepo)
	userHandler := userhalder.NewHandler(userSrv)
	cartRepo := cartrepo.NewRepository(db)
	cartSrv := cartsrv.NewService(cartRepo)
	cartHandler := cartHalder.NewHandler(cartSrv)
	productRepo := productrepo.NewRepository(db)
	productSrv := productsrv.NewService(productRepo)
	productHandler := producthalder.NewHandler(productSrv)
	s := apiserver.New(config)
	s.ConfigureRouter(productHandler, cartHandler, userHandler)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
