package main

import (
	"fmt"
	"sort"
)

func main() {
	var treasureCoordinates []string

	rowA := []string{"#", "#", "#", "#", "#", "#", "#", "#"}
	rowB := []string{"#", ".", ".", ".", ".", ".", ".", "#"}
	rowC := []string{"#", ".", "#", "#", "#", ".", ".", "#"}
	rowD := []string{"#", ".", ".", ".", "#", ".", "#", "#"}
	rowE := []string{"#", "X", "#", ".", ".", ".", ".", "#"}
	rowF := []string{"#", "#", "#", "#", "#", "#", "#", "#"}
	land := [][]string{rowA, rowB, rowC, rowD, rowE, rowF}
	printLand(land)

	land, treasureCoordinates = findTreasures(land, treasureCoordinates, 4, 1)
	printLand(land)

	land, treasureCoordinates = findTreasures(land, treasureCoordinates, 3, 1)
	printLand(land)

	land, treasureCoordinates = findTreasures(land, treasureCoordinates, 2, 1)
	printLand(land)

	land, treasureCoordinates = findTreasures(land, treasureCoordinates, 1, 1)
	printLand(land)

	sort.Strings(treasureCoordinates)
	fmt.Println("=== treasureCoordinates ===")
	fmt.Println(treasureCoordinates)
	fmt.Println("=== treasureCoordinates ===")
}

func findTreasures(land [][]string, treasureCoordinates []string, x int, y int) ([][]string, []string) {
	var columnsToFind []int

	for {
		if land[x+1][y] == "." {
			columnsToFind = append(columnsToFind, y)
		}

		if land[x][y] == "." {
			land[x][y] = "$"
			treasureCoordinate := fmt.Sprintf("(%d, %d)", x, y)
			treasureCoordinates = append(treasureCoordinates, treasureCoordinate)
		} else if land[x][y] == "#" {
			y = y - 2
			break
		}
		y++
	}

	x++
	for _, column := range columnsToFind {
		x2 := x
		for x2 < 8 {
			if land[x2][column] == "." {
				land[x2][column] = "$"
				treasureCoordinate := fmt.Sprintf("(%d, %d)", x2, column)
				treasureCoordinates = append(treasureCoordinates, treasureCoordinate)
			} else if land[x2][column] == "#" {
				break
			}
			x2++
		}
	}

	return land, treasureCoordinates
}

func printLand(land [][]string) {
	totalColumn := 8

	for i := 0; i < len(land); i++ {
		for j := 0; j < totalColumn; j++ {
			fmt.Print(land[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}