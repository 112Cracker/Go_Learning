package main

import "fmt"

func main() {
	line := []string{}
	plane := [][]string{}

	fmt.Printf("Use %T to make a line\n", line)
	fmt.Printf("Uss %T to make a plane\n", plane)

	//var line_var []string
	//var plane_var [][]string

	//line_make := make([][]string, 5)
	//plane_make := make([][]string, 5)

	//var line_var_array [5]string // fixed size array
	//var plane_var_array [5][5]string // fixed size multi-dimensional array
	fmt.Println("----------------------------------------------")

	//make a 5 by 5 game board
	//initialize the board cells to all ones
	game_board := make([][]int, 0)
	for i := 0; i < 5; i++ {
		row := make([]int, 0)
		for j := 0; j < 5; j++ {
			row = append(row, 1)
		}
		game_board = append(game_board, row)
	}
	fmt.Println("A new game board: ")
	for i:= 0; i < 5; i++ {
		fmt.Println(game_board[i])
	}

}