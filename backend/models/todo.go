package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/olucvolkan/todoApp/config"
)

type Todo struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateTodoRequest struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateTodoRequest struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type DeleteTodoRequest struct {
	ID uint        `json:"id"  binding:"required"`
}

func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(todo *Todo) (err error) {

	if err = config.DB.Delete(&Todo{}, todo.ID).Error; err != nil {
		return err
	}
	return nil
}
func UpdateATodo(todo *Todo) (err error) {

	if err = config.DB.Model(&todo).Where("id = ?", todo.ID).Updates(Todo{ Description: todo.Description, Status: todo.Status}).Error; err != nil {
		return err
	}
	return nil
}

func (b *Todo) TableName() string {
	return "todo"
}
