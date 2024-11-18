package controller

import (
	"eco-journal/entities"
	"eco-journal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JournalController struct {
	journalService service.JournalServiceInterface
}

func NewJournalController(journalService service.JournalServiceInterface) *JournalController {
	return &JournalController{journalService}
}

func (jc *JournalController) Create(c *gin.Context) {
	var journal entities.Journal
	if err := c.ShouldBindJSON(&journal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := jc.journalService.GetAuthorByID(journal.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	category, err := jc.journalService.GetCategoryByID(journal.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	createdJournal, err := jc.journalService.Create(&journal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdJournal.Author = *author
	createdJournal.Category = *category

	c.JSON(http.StatusCreated, gin.H{
		"message": "Journal added successfully",
		"data":    createdJournal,
	})
}

func (jc *JournalController) Update(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var journal entities.Journal
	if err := c.ShouldBindJSON(&journal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingJournal, err := jc.journalService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Journal not found"})
		return
	}

	journal.ID = existingJournal.ID

	author, err := jc.journalService.GetAuthorByID(journal.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	category, err := jc.journalService.GetCategoryByID(journal.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	updatedJournal, err := jc.journalService.Update(&journal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedJournal.Author = *author
	updatedJournal.Category = *category

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal update successfully",
		"data":    updatedJournal,
	})
}

func (jc *JournalController) Delete(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	journal, err := jc.journalService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Journal not found"})
		return
	}

	err = jc.journalService.Delete(journal.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Journal deleted successfully"})
}

func (jc *JournalController) GetAll(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	journals, err := jc.journalService.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems, err := jc.journalService.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data": journals,
	})
}

func (jc *JournalController) GetDetails(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	journal, err := jc.journalService.FindByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Journal not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": journal})
}
