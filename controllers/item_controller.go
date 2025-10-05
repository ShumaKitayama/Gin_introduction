package controllers

import (
	"gin-introduction/services"
	"net/http"
	"strconv"

	"gin-introduction/dto"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}	
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ItemController) FindById(ctx *gin.Context) {
	itemId,err := strconv.ParseUint(ctx.Param("id"),10,32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	item, err := c.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (c *ItemController) Create(ctx *gin.Context){
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	neweItem, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": neweItem})
}

func (c *ItemController) Update(ctx *gin.Context){
	itemId,err := strconv.ParseUint(ctx.Param("id"),10,32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedItem, err := c.service.Update(uint(itemId), input)
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"date": updatedItem})
}

func (c *ItemController) Delete(ctx *gin.Context) {
	itemId,err := strconv.ParseUint(ctx.Param("id"),10,32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.service.Delete(uint(itemId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}