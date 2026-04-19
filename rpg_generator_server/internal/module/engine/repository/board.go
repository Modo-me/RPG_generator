package repository

type Board struct {
	width  int
	height int
	cells  [][]string
}

func NewBoard(width, height int) *Board {
	cells := make([][]string, height)
	for i := range cells {
		cells[i] = make([]string, width)
	}
	return &Board{
		width:  width,
		height: height,
		cells:  cells,
	}
}
