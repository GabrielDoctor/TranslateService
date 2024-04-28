package handlers

import (
	translate "backend/services/translate"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Text string `json:"text"`
}

func TranslateText(c *gin.Context) {
	reader := io.Reader(c.Request.Body)
	input, err := io.ReadAll(reader)
	if err != nil {
		c.JSON(500, err)
		return
	}

	result := translate.TranslateText(string(input))

	var response Response
	response.Text = result
	c.JSON(200, response)
}

func TranslateChapter(c *gin.Context) {
	chapterId := c.Param("id")
	fmt.Println(chapterId)
	result, err := translate.TranslateChapter(chapterId)
	//log.Println(result)
	if err != nil {
		c.JSON(500, err)
		log.Println(err)
		return
	}
	var response Response
	response.Text = result
	//os.WriteFile("result.txt", []byte(result), 0644)
	c.JSON(200, response)
}
