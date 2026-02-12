package repository

import (
	"belajar-go-database-mysql/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

// CommentRepositoryImpl adalah implementasi dari CommentRepository yang menggunakan database SQL untuk menyimpan dan mengambil data komentar.
type CommentRepositoryImpl struct {
	DB *sql.DB
}

// NewCommentRepository membuat instance baru dari CommentRepositoryImpl dengan koneksi database yang diberikan.
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments (email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = id
	return comment, nil

}

func (repository CommentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	comment := entity.Comment{}

	rows, err := repository.DB.QueryContext(ctx, query, id)
	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		// jika ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// jika tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments "
	comment := entity.Comment{}

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		// jika ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil

}
