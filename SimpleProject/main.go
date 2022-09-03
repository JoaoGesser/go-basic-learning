package main

import "fmt"

type SimpleInformations struct {
	Name               string
	Age                int
	NumberOfIterations int
}

func main() {

	simple := SimpleInformations{Name: "Simple", Age: 90, NumberOfIterations: 2}

	if isSimple(simple) {
		defer interactWithSimple(simple)
	}

	fmt.Println("Starting...")

}

func interactWithSimple(simple SimpleInformations) {
	for i := 0; i < simple.NumberOfIterations; i++ {
		fmt.Println(simple)
	}
}

func isSimple(simple SimpleInformations) bool {
	if simple.Name == "Not Simple" {
		panic("It's not simple")
	}
	return true
}
