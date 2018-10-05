package main

import (
	"golang-study/gotour/ioreader/exercises"

	"golang.org/x/tour/reader"
)

func main() {
	reader.Validate(exercises.MyReader{})
	exercises.ExecuteRotReader()
}
