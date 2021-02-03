package main

import (
	"fmt"
	"github.com/ArtemZar/algorithms/lesson4/board"
)

func main() {

	board := board.InitBoard(5, 5, board.Coordinate{Row: 1, Col: 1})
	board.Main()
	fmt.Println("Board:")
	for key, _ := range board.Body {
		fmt.Printf("Координаты ячейки %v\n ", key)

	}
	for _, val := range board.Steps{
		fmt.Printf("Позиция и возможные варианты следующих ходов: %v\n", val)
	}


}

