package main

import (
	"rpg_generator/internal/database"
	"rpg_generator/internal/module/generation/handler"
	"rpg_generator/internal/module/generation/repository"
	"rpg_generator/internal/module/generation/service"
	"rpg_generator/internal/router"
)

func main() {
	db := database.DB_INIT()
	storyDao := repository.NewStoryDaoInstance(db)
	storyService := service.NewStoryService(storyDao)
	postHandler := handler.NewStoryHandler(storyService)

	routers := router.SetUpRouters(&router.Handlers{
		Story: postHandler,
	})
	routers.Run(":8080")

}
