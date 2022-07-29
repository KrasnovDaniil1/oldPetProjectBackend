package handlers

import (
	"net/http"
	"saveimage/storage"
	"github.com/gin-gonic/gin"
)

var params struct {
	search string
	limit  string
	offset string
}

func GetCards(c *gin.Context) {
	params.search = c.Query("search")
	params.limit = c.Query("limit")
	params.offset = c.Query("offset")

	if params.search == "" {
		allTableCard, err := storage.STORAGE.GetAllCard(params.limit, params.offset)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не получилось взять карточки из базы данных."})
			return

		}
		c.IndentedJSON(http.StatusOK, gin.H{"card": allTableCard})
	} else {
		allSearchCard, err := storage.STORAGE.GetSearchCard(params.search, params.limit, params.offset)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не получилось найти карточки."})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"card": allSearchCard})
	}
}
