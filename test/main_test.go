package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olucvolkan/todoApp/api/config"
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/api/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var err error

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createTestEnvironment() *gin.Engine {

	config.DB, err = gorm.Open("mysql", config.DbURL())

	if err != nil {
		fmt.Println("status: ", err)
	}
	config.DB.AutoMigrate(&models.Todo{})

	router := routes.SetupRouter()

	return router
}

func TestTodoCRUD(t *testing.T) {

	router := createTestEnvironment()


	t.Run("Create Todo endpoint", func(t *testing.T) {
		
		todo, _ := json.Marshal(models.CreateTodoRequest{Description: "Test Record", Status: "Todo"})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/create-todo/", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusCreated, w.Code)
	})


	t.Run("List Todo endpoint", func(t *testing.T) {
		w := performRequest(router, "GET", "/todo-list/")
		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
		
	})

	t.Run("Update Todo endpoint", func(t *testing.T) {

		todo, _ := json.Marshal(models.UpdateTodoRequest{ID:1,Description: "Test Record2", Status: "Todo"})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/update-todo/", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Delete Todo endpoint", func(t *testing.T) {

		todo, _ := json.Marshal(models.DeleteTodoRequest{ID: 1})

		fmt.Println(string(todo))
		req, err := http.NewRequest("POST", "/delete-todo/", bytes.NewReader(todo))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
