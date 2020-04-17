package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
    islandsCount := 0

    if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
    	fmt.Println("Grid empty...")
    	return islandsCount
    }
    fmt.Println("lengths: ", len(grid), len(grid[0]))
    for i := 0; i < len(grid); i++ {
    	for j := 0; j < len(grid[0]); j++ {
    		if grid[i][j] == '1' {
    			islandsCount++
    			fmt.Println("element is : ", grid[i][j])
    			//  passing arrays as slice will automagically does the copy-by-reference
    			traverse(grid[:], i, j)
    		}
    		fmt.Println("i and j", i, j)
    	}
    }
    fmt.Println("Grid after impact: %v", grid)
    return islandsCount
}

func traverse(grid [][]byte, x, y int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		traverse(grid[:], x+1, y)
		traverse(grid[:], x-1, y)
		traverse(grid[:], x, y+1)
		traverse(grid[:], x, y-1)
	} 
}

func main() {
	grid := [][]byte {{'1','0','0','1'}, {'1','0','1','0'}}
	fmt.Println("Islands count : ", numIslands(grid))
}