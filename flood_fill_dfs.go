package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	visited := map[string]bool{}
	replaceColor(image, sr, sc, oldColor, newColor, visited)

	return image
}

func replaceColor(
	image [][]int,
	row int,
	col int,
	oldColor int,
	newColor int,
	visited map[string]bool) {
	rowInBounds := 0 <= row && row < len(image)
	colInBounds := 0 <= col && col < len(image[0])
	if !rowInBounds || !colInBounds {
		return
	}

	if image[row][col] != oldColor {
		return
	}

	pos := fmt.Sprintf("%d,%d", row, col)
	fmt.Println("current position:", pos)
	if _, found := visited[pos]; found {
		return
	}

	visited[pos] = true
	image[row][col] = newColor

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, direction := range directions {
		rowOffset, colOffset := direction[0], direction[1]
		neighborRow, neighborCol := row+rowOffset, col+colOffset
		replaceColor(image, neighborRow, neighborCol, oldColor, newColor, visited)
	}
}

func main() {
	image := [][]int{
		{1, 1, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1},
	}

	sr := 2
	sc := 2
	newColor := 2

	fmt.Println("flood fill result", floodFill(image, sr, sc, newColor))
}

/*
output:
current position: 2,2
current position: 1,2
current position: 0,2
current position: 0,1
current position: 0,0
current position: 1,0
current position: 2,0
current position: 2,1
current position: 3,1
current position: 3,0
current position: 1,3
flood fill result [[2 2 2 0 1] [2 0 2 2 0] [2 2 2 0 1] [2 2 0 1 1]]
*/
