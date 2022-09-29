package server

import (
	"belajar/efishery/configs"
	"belajar/efishery/router"
	"fmt"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
)


func StartServer(){
	// create config and logger
	config, err := configs.New()
	if err != nil{

	}
	addr := fmt.Sprintf("%s:%v",config.Server.Host,config.Server.Port)
	conf := g.Config{
		Handler:             router.Router(config),
		ListenAddr:          addr,
		OnStarting: func() {
			log.Info().Msg("Service running at "+addr)
		},
	}

	g.Run(conf)
}
