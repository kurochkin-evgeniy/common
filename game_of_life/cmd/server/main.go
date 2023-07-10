package main

import (
	"game_of_life/package/gameoflife"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var GameFieldsMap map[string]*gameoflife.GameField = make(map[string]*gameoflife.GameField)

func getSessionHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	field, has := GameFieldsMap[id]
	if !has {
		field = gameoflife.GenRandomField(id, 100, 100, 10)
		GameFieldsMap[id] = field
	} else {
		field.Iterate()
	}

	w := len(field.Field)
	h := len(field.Field[0])

	var sb strings.Builder
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if field.Field[x][y][field.CurrentIndex] {
				sb.WriteString("#")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}

	text := sb.String()
	c.String(http.StatusOK, text)
}

func main() {
	httpEngine := gin.Default()

	httpEngine.GET("/:id", getSessionHandler)

	httpEngine.Run(":5008")
}
