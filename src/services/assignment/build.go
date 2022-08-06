package assignment

import "ggclass_log_service/src/config"

func BuildService() *service {
	repository := NewRepository(config.GetConfig().Mongo)
	return NewService(repository)
}
