package main

import (
	"ggclass_log_service/src/cmd"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/logger"
	"log"
)

func main() {

	config.Load()

	logger.InitLog()
	defer logger.SyncLog()

	rootCmd := cmd.GetRoot()

	err := rootCmd.Execute()

	if err != nil {
		log.Fatalln(err)
	}
}
