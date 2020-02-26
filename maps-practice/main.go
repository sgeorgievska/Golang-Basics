package main

import "fmt"

func main() {
	dogBreeds := map[string]int{
		"Maltese":       1,
		"Cavaliar King": 2,
		"Greyhound":     3,
		"Iggy":          4,
		"Shih tzu":      5,
	}
	fmt.Println(dogBreeds)

	//addition
	dogBreeds["German Sheppard"] = 6
	fmt.Println(dogBreeds)

	//delete
	delete(dogBreeds, "Iggy")
	fmt.Println(dogBreeds)

	//reading
	fmt.Println(dogBreeds["Maltese"])

	//blank identifier _ provides way to ignore left-hand side values in assignment
	_, ok := dogBreeds["Iggy"]
	fmt.Println(ok)

	fmt.Println(len(dogBreeds))
}
