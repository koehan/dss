package cmd

import (
	"dss/common/log"
	"dss/common/mongo"
	"dss/common/redis"
	"dss/core/config"
	"dss/core/global"
	"dss/core/router"
	"dss/core/scan"
	"dss/core/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func Consumer() *cobra.Command {
	var (
		cfg string
	)
	cmdConsumer := &cobra.Command{
		Use:   "consumer",
		Short: "Start Run Port Scan Consumer",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg) == 0 {
				_ = cmd.Help()
				os.Exit(0)
			}
			config.Init(cfg)
			conf := config.CoreConf
			log.Init(&conf.Log)
			if err := redis.Init(&conf.Redis); err != nil {
				log.Fatal("Init redis failed", zap.Error(err))
			}
			if err := mongo.Init(&conf.Mongo); err != nil {
				log.Fatal("Init mongo failed", zap.Error(err))
			}
			scan.Init(conf.Consumer.MaxWorkers, conf.Consumer.MaxQueue, log.Logger())
			gin.SetMode(conf.Consumer.Mode)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			if err := server.Run(router.NewHttpRouter(), global.Consumer); nil != err {
				log.Error("server run error", zap.Error(err))
			}
		},
	}
	cmdConsumer.Flags().StringVarP(&cfg, "conf", "c", "", "server config [toml]")
	return cmdConsumer
}
