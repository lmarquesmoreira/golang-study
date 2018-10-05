package main

import "golang.org/x/tour/reader"

func main() {
	reader.Validate(MyReader{})

	executeRotReader()
}
