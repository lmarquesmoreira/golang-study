package pseudoheranca

import "fmt"

type Car struct {
	name string
	currentSpeed int
}

type Ferrari struct {
	Car
	turboOn bool
}

func ExecuteDemoPseudoHeranca(){

	f:= Ferrari{}
	f.name = "f40"
	f.currentSpeed = 0
	f.turboOn = false

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}
