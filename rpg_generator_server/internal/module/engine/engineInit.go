package engine

import "rpg_generator/internal/module/engine/repository"

var GlobalBoard *repository.Board

func InitBoard(width, height int) {
	GlobalBoard = repository.NewBoard(width, height)
}
