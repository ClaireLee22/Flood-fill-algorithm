package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}

	visited := map[string]bool{getPosString(sr, sc): true}
	queue := [][]int{{sr, sc}}
	for len(queue) > 0 {
		fmt.Println("queue:", queue)
		currentRow, currentCol := queue[0][0], queue[0][1]
		queue = queue[1:]
		image[currentRow][currentCol] = newColor

		directions := [][]int{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1},
		}

		for _, direction := range directions {
			rowOffset, colOffset := direction[0], direction[1]
			neighborRow := currentRow + rowOffset
			neighborCol := currentCol + colOffset
			pos := getPosString(neighborRow, neighborCol)
			if isInBounds(image, neighborRow, neighborCol) && image[neighborRow][neighborCol] == oldColor {
				if _, found := visited[pos]; found {
					continue
				}
				visited[pos] = true
				queue = append(queue, []int{neighborRow, neighborCol})
			}

		}
	}
	return image
}

func isInBounds(image [][]int, row int, col int) bool {
	rowInBounds := 0 <= row && row < len(image)
	colInBounds := 0 <= col && col < len(image[0])
	return rowInBounds && colInBounds
}

func getPosString(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
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

	fmt.Println("flood fill result:", floodFill(image, sr, sc, newColor))
}

/*
output:
queue: [[2 2]]
queue: [[1 2] [2 1]]
queue: [[2 1] [0 2] [1 3]]
queue: [[0 2] [1 3] [3 1] [2 0]]
queue: [[1 3] [3 1] [2 0] [0 1]]
queue: [[3 1] [2 0] [0 1]]
queue: [[2 0] [0 1] [3 0]]
queue: [[0 1] [3 0] [1 0]]
queue: [[3 0] [1 0] [0 0]]
queue: [[1 0] [0 0]]
queue: [[0 0]]
flood fill result: [[2 2 2 0 1] [2 0 2 2 0] [2 2 2 0 1] [2 2 0 1 1]]
*/
