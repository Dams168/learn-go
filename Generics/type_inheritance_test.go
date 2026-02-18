package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Employee adalah interface dasar.
// Semua tipe yang punya method GetName() otomatis dianggap sebagai Employee.
type Employee interface {
	GetName() string
}

// T Employee artinya:
// T harus bertipe apapun ASAL mengimplementasikan interface Employee.
// Jadi T boleh Manager, VicePresident, dll
// selama punya method GetName().
func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

// Manager "mewarisi" Employee secara tidak langsung,
// karena dia juga punya GetName().
//
// Di Go tidak ada keyword "extends",
// tapi jika interface memiliki method yang sama,
// maka dia kompatibel.
type Manager interface {
	GetName() string
	GetManagerName() string
}

// Struct konkret yang mengimplementasikan Manager
type MyManager struct {
	Name string
}

// Implementasi method GetName()
// Ini membuat MyManager memenuhi interface Employee juga.
func (m *MyManager) GetName() string {
	return m.Name
}

// Implementasi method khusus Manager
func (m *MyManager) GetManagerName() string {
	return m.Name
}

// Interface lain yang juga punya GetName()
type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

// Struct konkret
type MyVicePresident struct {
	Name string
}

// Implementasi GetName()
// Karena ada GetName(), dia otomatis Employee juga.
func (m *MyVicePresident) GetName() string {
	return m.Name
}

// Method khusus VicePresident
func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {

	// Di sini kita memanggil generic function:
	// GetName[T Employee]
	//
	// Manager punya GetName(), jadi valid.
	assert.Equal(t, "Eko", GetName[Manager](&MyManager{Name: "Eko"}))

	// VicePresident juga punya GetName(), jadi valid.
	assert.Equal(t, "Eko", GetName[VicePresident](&MyVicePresident{Name: "Eko"}))
}

// di Go tidak ada konsep inheritance seperti di bahasa lain, tapi kita bisa menggunakan interface untuk mencapai efek yang mirip.
// berikut adalah contoh bagaimana kita bisa menggunakan interface untuk membuat hierarki tipe yang mirip dengan inheritance:
// Employee (parent)
//    ↓
// Manager (child)
//    ↓
// VicePresident (child)
