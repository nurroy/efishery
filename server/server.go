package server

import (
	"belajar/efishery/configs"
	"belajar/efishery/router"
	"fmt"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func StartServer(){
	addr := fmt.Sprintf("%s:%v",configs.Config.ServiceHost,configs.Config.ServicePort)
	conf := g.Config{
		Handler:             router.Router(),
		ListenAddr:          addr,
		OnStarting: func() {
			log.Info().Msg("Service running at "+addr)
		},
	}

	g.Run(conf)
}
