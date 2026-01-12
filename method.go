package main

import "fmt"

type student struct {
	Name string
	Age  int
}

// method adalah function yang di deklarasikan dengan receiver tipe data tertentu
// receiver adalah tipe data yang menerima method tersebut
// contoh method dengan receiver bertipe struct student
func (s student) sayName() {
	fmt.Println("My name", s.Name)
}

func (s student) sayAge() int {
	return s.Age
}

// method dengan receiver value
// perubahan pada receiver tidak akan mempengaruhi data asli
// karena yang diubah adalah salinan dari data asli
// receiver value ditandai tanpa tanda *
func (s student) changeName(newName string) {
	s.Name = newName
}

// method dengan receiver pointer
// perubahan pada receiver akan mempengaruhi data asli
// karena yang diubah adalah alamat memory dari data asli
// receiver pointer ditandai dengan tanda *
func (s *student) changeName2(newName string) {
	s.Name = newName
}

func main() {
	s1 := student{Name: "john", Age: 20}

	s1.sayName()
	fmt.Println("My age", s1.sayAge())

	s1.changeName("doe")
	fmt.Println("After change name", s1.Name)
	s1.changeName2("dodo")
	fmt.Println("After change name 2", s1.Name)
}
