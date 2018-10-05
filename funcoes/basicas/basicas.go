package main

import (
	"fmt"
)

func f1() {
	fmt.Println("f1()")
}

func f2(p1 string, p2 string) {
	fmt.Printf("f2: %s %s\n", p1, p2)
}

func f3() string {
	return "f3()"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("f4: %s %s", p1, p2)
}

func f41(p1, p2 string, i int) string {
	return fmt.Sprintf("f41: %s %s %d", p1, p2, i)
}

func f5() (string, string) {
	return "return 1", "return 2"
}
func main() {
	f1()
	f2("ola", "mundo")

	r3, r4, r41 := f3(), f4("novo", "mundo"), f41("mundo", "novo", 91)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r41)

	r51, r52 := f5()
	fmt.Println(r51, r52)
}
