package main

import "fmt"

//struct adalah tipe data buatan sendiri yang dapat mengelompokkan beberapa variabel dengan tipe data berbeda menjadi satu
// sederhananya, struct adalah kumpulan dari beberapa field/properti
// mirip seperti class di bahasa pemrograman lain, namun struct tidak memiliki method

type Person struct {
	Name string
	Age  int
}

// embedding struct
type Student struct {
	grade int
	Person
}

func main() {

	// membuat instance dari struct Person
	var p1 Person
	p1.Name = "Alice"
	p1.Age = 30

	// membuat instance lain dari struct Person dengan cara lain
	var p2 = Person{Name: "Bob", Age: 25}
	var p3 = Person{"Charlie", 28} // tanpa menyebutkan nama field

	// membuat instance dengan hanya mengisi sebagian field, lainnya akan bernilai default
	var p4 = Person{Name: "Diana"} // property bisa tidak berurutan

	// membuat instance dari struct Student yang meng-embed struct Person
	var s1 = Student{}
	s1.Name = "Eve"
	s1.Age = 20
	s1.grade = 3

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)

	fmt.Println("")
	fmt.Println(p2)

	fmt.Println("")
	fmt.Println(p3)

	fmt.Println("")
	fmt.Println(p4)

	fmt.Println("")
	fmt.Println(s1)
	fmt.Println(s1.Name)
	fmt.Println(s1.grade)
	fmt.Println(s1.Age)
	fmt.Println(s1.Person.Name)

}
