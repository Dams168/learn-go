package repository

import (
	belajar_go_database_mysql "belajar-go-database-mysql"
	"belajar-go-database-mysql/entity"
	"context"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_go_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "testrepo@mail.com",
		Comment: "Ini Test Comment Repo",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_go_database_mysql.GetConnection())

	ctx := context.Background()

	result, err := CommentRepository.FindById(ctx, 10)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_go_database_mysql.GetConnection())

	ctx := context.Background()

	result, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
