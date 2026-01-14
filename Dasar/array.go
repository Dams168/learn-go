package main

import "fmt"

func arrayExample() {
	fruits := [4]string{"appple", "mango", "cherry", "Melon"}
	Even := [...]int{2, 4, 6, 8}

	fmt.Println(fruits)
	fmt.Println(len(Even))

	for i := range len(Even) {
		fmt.Print(Even[i], " ")
	}

	for _, fruit := range fruits {
		fmt.Print(fruit)
	} //gunakan variabel penggangguran( _ ), karena variabel i tidak digunakan

}

func sliceExample() {
	//mirip array, bedanya tidak perlu deklarasi jumlah element
	//jika tidak dideskripsikan itu berarti slice
	vegetables := []string{"Brocoli", "Onion", "Garlic"}

	veg := vegetables[0:2]

	// fmt.Println(vegetables)
	// fmt.Println(vegetables[0:2])
	// // fmt.Println(len(vegetables))
	fmt.Println(len(veg))
	fmt.Println(cap(veg))

	var aVeg = append(veg, "tomato")

	fmt.Println(aVeg)
	fmt.Println(veg)
	fmt.Println(vegetables)
}
