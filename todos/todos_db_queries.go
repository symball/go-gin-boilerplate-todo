package todos

import (
	"github.com/gin-gonic/gin"
	"github.com/symball/go-gin-boilerplate/storage"
)

func Add(todo *Todo, ctx *gin.Context) (*Todo, error) {
	_, err := storage.DBGet().NewInsert().Model(todo).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Retrieve a single Todo record from the DB
func GetOneById(TodoId int64, ctx *gin.Context) (*Todo, error) {
	todo := new(Todo)
	err := storage.DBGet().NewSelect().Model(todo).Where("id = ?", TodoId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Retrieve all Todo records from the DB
func GetAll(ctx *gin.Context) ([]Todo, error) {
	var todos []Todo
	// TODO Look more in to Bun in order to use the type for selected columns. Believe it is just extending model but, for time to implement, better value elsewhere
	err := storage.DBGet().NewSelect().Model(&todos).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func Update(todo *Todo, ctx *gin.Context) (*Todo, error) {
	_, err := storage.DBGet().NewUpdate().Model(todo).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
