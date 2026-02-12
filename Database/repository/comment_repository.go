package repository

import (
	"belajar-go-database-mysql/entity"
	"context"
)

// CommentRepository adalah interface yang mendefinisikan metode-metode untuk menyimpan dan mengambil data komentar dari database.
// Ini adalah contract yang harus diimplementasikan oleh setiap repository komentar.
type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int64) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
