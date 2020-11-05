package service

import (
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/api/schema"
	"strconv"
)

func CreateTodoListSchema(todos []models.Todo) schema.TodoListSchema {

	var listSchema = schema.TodoListSchema{
		Todo:       schema.Todo{Title: "todo"},
		Inprogress: schema.InProgress{Title: "in-progress"},
		Done:       schema.Done{Title: "done"},
	}
	for _, todo := range todos {
		listSchema = AddItem(listSchema, todo.ID, todo.Description, todo.Status)
	}

	return listSchema
}
func AddItem(todoListSchema schema.TodoListSchema, todoId uint, todoText string, todoStatus string) schema.TodoListSchema {

	var todoIdString = strconv.FormatUint(uint64(todoId), 10)

	item := schema.Item{ID: todoIdString, Description: todoText, Status: todoStatus}
	todoListSchema.Todo.Items = make([]schema.Item, 0)
	todoListSchema.Inprogress.Items = make([]schema.Item, 0)
	todoListSchema.Done.Items = make([]schema.Item, 0)
	switch todoStatus {
	case "todo":
		todoListSchema.Todo.Items = append(todoListSchema.Todo.Items, item)
	case "in-progress":
		todoListSchema.Inprogress.Items = append(todoListSchema.Inprogress.Items, item)
	case "done":
		todoListSchema.Done.Items = append(todoListSchema.Done.Items, item)
	}

	return todoListSchema

}
