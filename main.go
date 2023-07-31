package main

import "fmt"

var numberA int = 4
var numberB *int = &numberA

func main() {
	fmt.Println("Number A (value)", numberA)
	fmt.Println("Number A (address)", &numberA)

	fmt.Println("Number B (value)", *numberB)
	fmt.Println("Number B (address)", numberB)

	fmt.Println()

	// ubah nilai variabel numberB
	var tes int = 5
	// numberB = &tes // sama saja mengganti alamat memori ke variabel tes
	*numberB = tes
	fmt.Println("Number A (value)", numberA)
	fmt.Println("Number A (address)", &numberA)

	fmt.Println("Number B (value)", *numberB)
	fmt.Println("Number B (address)", numberB)

	fmt.Println()

	// ubah variabel numberA
	numberA = 6
	fmt.Println("Number A (value)", numberA)
	fmt.Println("Number A (address)", &numberA)

	fmt.Println("Number B (value)", *numberB)
	fmt.Println("Number B (address)", numberB)

	fmt.Println()
	// pointer sebagai parameter di function
	var number int = 4
	fmt.Println("before", number)

	change(&number, *numberB)
	fmt.Println("after", number)
}

// mendeklarasikan function dimana parameter pertama bertipe pointer int
func change(original *int, number int) {

	// parameter pertma diambil nilai aslinya dengan menggunakan metode dereferencing
	*original = number
}
