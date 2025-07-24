package main

import (
	"erzi_new/internal/apiserver"
	cartitemhalder "erzi_new/internal/handler/cartItem"
	producthalder "erzi_new/internal/handler/product"
	userhalder "erzi_new/internal/handler/user"
	cartrepo "erzi_new/internal/repository/cart"
	cartitemrepo "erzi_new/internal/repository/cartItem"
	productrepo "erzi_new/internal/repository/product"
	userrepo "erzi_new/internal/repository/user"
	cartsrv "erzi_new/internal/service/cart"
	cartitemsrv "erzi_new/internal/service/cartItem"
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
	//repo
	cartRepo := cartrepo.NewRepository(db)
	productRepo := productrepo.NewRepository(db)
	userRepo := userrepo.NewRepository(db)
	cartItemRepo := cartitemrepo.NewRepository(db)
	//srv
	cartSrv := cartsrv.NewService(cartRepo)
	productSrv := productsrv.NewService(productRepo)
	userSrv := usersrv.NewService(userRepo, cartRepo)
	cartItemSrv := cartitemsrv.NewService(cartItemRepo, cartRepo)
	//handler
	userHandler := userhalder.NewHandler(userSrv)
	cartItemHandler := cartitemhalder.NewHandler(cartItemSrv, cartSrv)
	productHandler := producthalder.NewHandler(productSrv)

	s := apiserver.New(config)
	s.ConfigureRouter(productHandler, userHandler, cartItemHandler)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
