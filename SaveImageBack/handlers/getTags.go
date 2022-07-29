package handlers

import (
	"net/http"
	"saveimage/storage"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	search := c.Query("search")
	if search == "" {
		allTableTags, err := storage.STORAGE.GetAllTags()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не получилось взять карточки из базы данных."})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"tags": allTableTags})
	} else {
		tableTags, err := storage.STORAGE.GetSearchTag(search)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не получилось взять карточки из базы данных."})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"tag": tableTags})
	}
}
