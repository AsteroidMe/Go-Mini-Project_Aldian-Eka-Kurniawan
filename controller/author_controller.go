package controller

import (
	"eco-journal/entities"
	"eco-journal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorService service.AuthorServiceInterface
}

func NewAuthorController(authorService service.AuthorServiceInterface) *AuthorController {
	return &AuthorController{authorService}
}

func (ac *AuthorController) Create(c *gin.Context) {
	var author entities.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAuthor, err := ac.authorService.Create(&author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Author added successfully",
		"data":    createdAuthor,
	})
}

func (ac *AuthorController) Update(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var author entities.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingAuthor, err := ac.authorService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	author.ID = existingAuthor.ID

	updatedAuthor, err := ac.authorService.Update(&author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author update successfully",
		"data":    updatedAuthor,
	})
}

func (ac *AuthorController) Delete(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := ac.authorService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	err = ac.authorService.Delete(author.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}

func (ac *AuthorController) GetAll(c *gin.Context) {
	pageStr := c.Param("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	authors, pagination, err := ac.authorService.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pagination": pagination,
		"data":       authors,
	})
}

func (ac *AuthorController) GetDetails(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := ac.authorService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}
