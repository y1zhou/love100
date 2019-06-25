package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y1zhou/love100/backend/db"
)

// CreateContent ...
func CreateContent(c *gin.Context) {
	var json db.ContentCreateForm
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": err.Error(),
		})
		return
	}
	item := db.Contents{
		Title:  json.Title,
		Status: json.Status,
	}
	if err := db.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "",
			"err": err,
		})
	} else {
		// Create new item
		c.JSON(http.StatusCreated, gin.H{
			"msg": item.ID,
			"err": "",
		})
	}
}

// DeleteContent ...
func DeleteContent(c *gin.Context) {
	var json db.ContentDeleteForm
	if err := c.ShouldBind(&json); err != nil {
		// Form validation errors
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": err.Error(),
		})
		return
	}
	var content db.Contents
	if err := db.DB.Unscoped().
		Where(json.ItemIDs).
		Delete(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Failed to delete item",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": json.ItemIDs,
			"err": "",
		})
	}
}

// UpdateContent ...
func UpdateContent(c *gin.Context) {
	var json db.ContentUpdateForm
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": err.Error(),
		})
		return
	}
	var contents db.Contents
	if err := db.DB.Where("title = ?", json.OldTitle).First(&contents).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "",
			"err": "Item ID not found in database.",
		})
		return
	}
	if err := db.DB.Model(&contents).
		Updates(map[string]interface{}{
			"Title":  json.NewTitle,
			"Status": json.Status,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Update items failed",
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": contents.ID,
			"err": "",
		})
	}
}

// FetchAllContents ...
func FetchAllContents(c *gin.Context) {
	var contents []db.Contents
	if err := db.DB.Find(&contents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Failed to get all items",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  len(contents),
			"err":  "",
			"data": contents,
		})
	}
}
