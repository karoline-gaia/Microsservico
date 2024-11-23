package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karoline-gaia/go-categories-mvc/internal/repositories"
	use_cases "github.com/karoline-gaia/go-categories-mvc/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase()
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"sucess": true})
}
