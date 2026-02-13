package repository

import (
	"belajar-go-restful-api/model/domain"
	"context"
	"database/sql"
)

// ini adalah contract / blueprint dari repository category
// fungsinya untuk mendefinisikan method-method yang harus diimplementasikan pada repository category
// kenapa harus ada contract? supaya repository lain yang mengimplementasikan contract ini harus memiliki method-method yang sama
// jika tidak memiliki method yang sama, maka akan terjadi error pada saat kompilasi
type CategoryRepository interface {
	// Parameter ctx digunakan untuk mengirimkan informasi tentang request yang sedang diproses, seperti deadline, cancelation signal, dan lain-lain (Context)
	// Parameter tx digunakan untuk mengirimkan informasi tentang transaksi yang sedang diproses, seperti commit, rollback, dan lain-lain (Database Transaction)
	// Parameter category digunakan untuk mengirimkan informasi tentang category yang sedang diproses, seperti id, name, dan lain-lain
	// Return value domain.Category digunakan untuk mengirimkan informasi tentang category yang sudah diproses, seperti id, name, dan lain-lain
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
}
