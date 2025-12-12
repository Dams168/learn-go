package main

import (
	"fmt"
)

func mapExample() {
	//map adalah tipe data asosiatif yang menyimpan pasangan kunci-nilai
	var person map[string]int
	person = map[string]int{} //inisialisasi

	person["age"] = 50
	person["phone"] = 1234

	fmt.Println("Age :", person["age"])
	fmt.Println("Phone :", person["phone"])

	data := map[string]int{}

	data["num"] = 60

	fmt.Println(data["num"])
	// chicken := make(map[string]int)

	// chicken["chick"] = 60
	// fmt.Println(chicken)

	numCodes := map[string]int{
		"Perfect": 100,
		"Awesome": 90,
		"Good":    80,
		"Enough":  70,
	}

	fmt.Println(numCodes)
	fmt.Println(numCodes["Good"])

	for key, num := range numCodes {
		fmt.Println(key, " : ", num)
	}
	fmt.Print(len(numCodes))

	delete(numCodes, "Enough")
	fmt.Println(numCodes)

	// slice and map combination
	students := []map[string]string{
		{"name": "Abdul", "gender": "male"},
		{"name": "Yaya", "gender": "female"},
		{"name": "John", "gender": "male"},
		{"name": "Asep", "gender": "male", "address": "Cipatat"},
	}

	for k, v := range students {
		// fmt.Println(v["name"], " : ", v["gender"], " : ", v["address"])
		fmt.Println(k, " : ", v["name"])
	}

}
