package main

import "golang.org/x/tour/pic"

func myPic(dx, dy int) (result [][]uint8) {

	result = make([][]uint8, dy)

	for x := range result {
		result[x] = make([]uint8, dx)

		for y := range result[x] {
			result[x][y] = uint8((x ^ y) / 2)
		}
	}

	return
}

func main() {
	pic.Show(myPic)
}
