package main

import (
	"rpg_generator/internal/database"
	"rpg_generator/internal/module/task/handler"
	"rpg_generator/internal/module/task/repository"
	"rpg_generator/internal/module/task/service"
	"rpg_generator/internal/router"
)

func main() {
	db := database.DB_INIT()
	taskDao := repository.NewTaskDaoInstance(db)
	taskResultDao := repository.NewTaskResultDaoInstance(db)
	taskService := service.NewTaskService(taskDao, taskResultDao)
	postHandler := handler.NewTaskHandler(taskService)

	routers := router.SetUpRouters(&router.Handlers{
		Task: postHandler,
	})
	routers.Run(":8080")

}
