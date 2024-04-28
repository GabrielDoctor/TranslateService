package repositorys

import (
	"backend/models"
	"log"
	"os"
)

func GetChapterContent(chapterId string) (string, error) {

	var content_path string
	err := models.DB.QueryRow("SELECT content_path FROM chapters WHERE id = $1", chapterId).Scan(&content_path)
	log.Println("path: ", content_path)
	if err != nil {
		return "", err
	}
	content, err := os.ReadFile(content_path)
	//fmt.Println(string(content))
	if err != nil {
		return "", err
	}
	return string(content), nil
}
