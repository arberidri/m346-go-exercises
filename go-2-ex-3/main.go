package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{

		104: "Modul 104: ",
		117: "Modul 117: ",
		346: "Modul 346: ",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[320] = "Modul 320:"
	// TODO: replace one
	modules[320] = "test"
	fmt.Println(modules)
}
