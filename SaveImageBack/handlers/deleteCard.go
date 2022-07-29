package handlers

import (
	"net/http"
	"saveimage/storage"
	"github.com/gin-gonic/gin"
)

func DeleteCard(c *gin.Context) {
	id := c.Param("id")
	err := storage.STORAGE.DeleteCard(id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": "Не получилось удалить карточку"})
		return

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Карточка удалена"})

}
