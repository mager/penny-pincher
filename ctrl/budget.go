package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mager/penny-pincher/entity"
)

func GetBudget(c *gin.Context) {
	var b entity.Budget
	// err := Models.GetAllTodos(&todo)

	// Fetch budget from DB
	c.JSON(http.StatusOK, b)
}
