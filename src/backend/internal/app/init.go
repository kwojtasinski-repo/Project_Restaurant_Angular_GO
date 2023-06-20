package app

import (
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/config"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/api"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
)

func InitApp(configFile config.Config) {
	api.InitObjectCreator(configFile)
	hashId, err := api.CreateHashId()
	if err != nil {
		panic(err)
	}

	dto.InitialIdObject(hashId)
}
