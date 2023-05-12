/*
 * Symbolic Boilerplate
 *
 * An API spec for managing todo lists
 *
 * API version: 1.0.0
 * Contact: contact@simonball.me
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package todos

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/symball/go-gin-boilerplate/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TodosGet - Get a listing of todo in system
func TodosGet(c *gin.Context) {
	todos, err := GetAll(c)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	todoListing := []TodoListing{}
	copier.Copy(&todoListing, &todos)
	c.JSON(http.StatusOK, TodosGet200Response{
		Message: "todosGetAll",
		Data:    todoListing,
	})
}

// TodosGetById - Get detailed information about a particular todo
func TodosGetById(c *gin.Context) {

	uri := TodosIdURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "TodoId required and a number"})
		return
	}

	todo, err := GetOneById(uri.TodoId, c)
	if err != nil {
		fmt.Printf("Error: %v", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, TodosGetById200Response{
		Message: "TodosGetById",
		Data:    *todo,
	})
}

// TodosPost - Create a new todo in system
func TodosPost(c *gin.Context) {

	var newTodo TodosPostRequest
	var todo Todo
	if errors := c.ShouldBindJSON(&newTodo); errors != nil {
		outPutData := lib.HandleValidationErrors(errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": outPutData})
	}

	copier.Copy(&todo, &newTodo)
	if _, err := Add(&todo, c); err != nil {
		fmt.Printf(err.Error())
	}

	c.JSON(http.StatusOK, TodosPost200Response{
		Message: "todosPost",
		Data:    todo,
	})
}

// TodosPutById - Update a particular todo
func TodosPutById(c *gin.Context) {

	uri := TodosIdURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "TodoId required and a number"})
		return
	}

	var putTodo TodosPutByIdRequest
	var todo Todo
	if errors := c.ShouldBindJSON(&putTodo); errors != nil {
		outPutData := lib.HandleValidationErrors(errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": outPutData})
	}
	copier.Copy(&todo, &putTodo)
	todo.Id = uri.TodoId
	Update(&todo, c)

	c.JSON(http.StatusOK, TodosPutById200Response{
		Message: "todosPut",
		Data:    todo,
	})
}